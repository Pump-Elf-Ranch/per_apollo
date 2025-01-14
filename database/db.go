package database

import (
	"context"
	"fmt"
	"github.com/Pump-Elf-Ranch/per_apollo/database/block_listener"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
	"github.com/Pump-Elf-Ranch/per_apollo/database/common"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"gorm.io/driver/postgres"
	gormDb "gorm.io/gorm"
	"path/filepath"

	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database/utils"
	_ "github.com/Pump-Elf-Ranch/per_apollo/database/utils/serializers"
	"github.com/Pump-Elf-Ranch/per_apollo/synchronizer/retry"
)

type DB struct {
	gorm           *gormDb.DB
	Blocks         common.BlocksDB
	ContractEvents common.ContractEventsDB
	CreateTable    common.CreateTableDB
	MintListed     business.MintListedDB
	BuyPropListed  business.BuyPropListedDB
	BlockListener  block_listener.BlockListenerDB
}

func NewDB(dbConfig config.Database) (*DB, error) {
	dbLog := log.New()

	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.DbHost, dbConfig.DbName)
	if dbConfig.DbPort != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.DbPort)
	}
	if dbConfig.DbUser != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.DbUser)
	}
	if dbConfig.DbPassword != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.DbPassword)
	}
	gormConfig := gormDb.Config{
		Logger:                 utils.NewLogger(dbLog),
		SkipDefaultTransaction: true,
		CreateBatchSize:        500,
	}
	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	gorm, err := retry.Do[*gormDb.DB](context.Background(), 10, retryStrategy, func() (*gormDb.DB, error) {
		gorm, err := gormDb.Open(postgres.Open(dsn), &gormConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
		}
		return gorm, nil
	})
	if err != nil {
		return nil, err
	}
	db := &DB{
		gorm:           gorm,
		Blocks:         common.NewBlocksDB(gorm),
		ContractEvents: common.NewContractEventsDB(gorm),
		CreateTable:    common.NewCreateTableDB(gorm),
		MintListed:     business.NewMintListedDB(gorm),
		BuyPropListed:  business.NewBuyPropListedDB(gorm),
		BlockListener:  block_listener.NewBlockListenerDB(gorm),
	}
	return db, nil
}

func (db *DB) Transaction(fn func(db *DB) error) error {
	return db.gorm.Transaction(func(gorm *gormDb.DB) error {
		txDB := &DB{
			gorm:           gorm,
			Blocks:         common.NewBlocksDB(gorm),
			ContractEvents: common.NewContractEventsDB(gorm),
			CreateTable:    common.NewCreateTableDB(gorm),
			MintListed:     business.NewMintListedDB(gorm),
			BuyPropListed:  business.NewBuyPropListedDB(gorm),
			BlockListener:  block_listener.NewBlockListenerDB(gorm),
		}
		return fn(txDB)
	})
}

func (db *DB) Close() error {
	sql, err := db.gorm.DB()
	if err != nil {
		return err
	}
	return sql.Close()
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Failed to process migration file: %s", path))
		}
		if info.IsDir() {
			return nil
		}
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return errors.Wrap(readErr, fmt.Sprintf("Error reading SQL file: %s", path))
		}
		// Check if migrations have already been run
		var migrationCount int
		err = db.gorm.Raw("SELECT COUNT(*) FROM migrations WHERE version = ?", path).Scan(&migrationCount).Error
		if err != nil {
			log.Error("failed to check migration status", "err", err)
			return err
		}
		if migrationCount > 0 {
			log.Info("migrations already run, skipping")
			return nil
		}
		execErr := db.gorm.Exec(string(fileContent)).Error
		if execErr != nil {
			return errors.Wrap(execErr, fmt.Sprintf("Error executing SQL script: %s", path))
		}

		// Record the migration as completed
		err = db.gorm.Exec("INSERT INTO migrations (version) VALUES (?)", path).Error
		if err != nil {
			log.Error("failed to record migration", "err", err)
			return err
		}
		return nil
	})
	return err
}
