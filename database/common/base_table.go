package common

import (
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"
)

type CreateTableDB interface {
	CreateTable(tableName, realTableName string)
	InitMigration()
}

type createTableDB struct {
	gorm *gorm.DB
}

func NewCreateTableDB(db *gorm.DB) CreateTableDB {
	return &createTableDB{gorm: db}
}

func (dao *createTableDB) CreateTable(tableName, realTableName string) {
	err := dao.gorm.Exec("CREATE TABLE IF NOT EXISTS " + tableName + "(like " + realTableName + " including all)").Error
	if err != nil {
		log.Error("create table from base table fail", "err", err)
	}
}

func (dao *createTableDB) InitMigration() {
	err := dao.gorm.Exec(`CREATE TABLE IF NOT EXISTS migrations (
        version VARCHAR NOT NULL,
        applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`).Error
	if err != nil {
		log.Error("failed to create migrations table", "err", err)
	}
}
