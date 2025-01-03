package main

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/urfave/cli/v2"

	PerSynchronizer "github.com/Pump-Elf-Ranch/per_apollo"
	"github.com/Pump-Elf-Ranch/per_apollo/common/cliapp"
	"github.com/Pump-Elf-Ranch/per_apollo/common/opio"
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/create_table"
)

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Value:   "./apollo.yaml",
		Aliases: []string{"c"},
		Usage:   "path to config file",
		EnvVars: []string{"APOLOLO_CONFIG"},
	}
	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path to migrations folder",
		EnvVars: []string{"RUNES_MIGRATIONS_DIR"},
	}
)

func runApiServer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("running api Server...")
	cfg, err := config.New(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return PerSynchronizer.NewApi(ctx.Context, cfg, shutdown)
}

func runSyncServer(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("running sync Server...")
	cfg, err := config.New(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return nil, err
	}
	return PerSynchronizer.NewSync(ctx.Context, cfg, shutdown)
}

func runMigrations(ctx *cli.Context) error {
	ctx.Context = opio.CancelOnInterrupt(ctx.Context)
	log.Info("running migrations...")
	cfg, err := config.New(ctx.String(ConfigFlag.Name))
	if err != nil {
		log.Error("failed to load config", "err", err)
		return err
	}
	db, err := database.NewDB(cfg.MasterDb)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	err = db.ExecuteSQLMigration(ctx.String(MigrationsFlag.Name))
	if err != nil {
		return err
	}
	for i := range cfg.RPCs {
		log.Info("create chain table by chainId", "chainId", cfg.RPCs[i].ChainId)
		create_table.CreateTableFromTemplate(strconv.Itoa(int(cfg.RPCs[i].ChainId)), db)
	}
	log.Info("running migrations and create table from template success")
	return nil
}

func newCli(GitCommit string, GitDate string) *cli.App {
	flags := []cli.Flag{ConfigFlag}
	migrationFlags := []cli.Flag{MigrationsFlag, ConfigFlag}
	return &cli.App{
		Version:              params.VersionWithCommit(GitCommit, GitDate),
		Description:          "An indexer of all optimism events with a serving api layer",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "api",
				Flags:       flags,
				Description: "Runs the runApiServer",
				Action:      cliapp.LifecycleCmd(runApiServer),
			},
			{
				Name:        "sync",
				Flags:       flags,
				Description: "Runs the runSyncServer",
				Action:      cliapp.LifecycleCmd(runSyncServer),
			},
			{
				Name:        "migrate",
				Flags:       migrationFlags,
				Description: "Runs the database migrations",
				Action:      runMigrations,
			},
			{
				Name:        "version",
				Description: "print version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
