package deposit_prop_worker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/tasks"
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
)

type WorkerProcessor struct {
	db    *database.DB
	tasks tasks.Group
	cfg   *config.Config
}

func NewWorkerProcessor(db *database.DB, shutdown context.CancelCauseFunc, cfg *config.Config) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db: db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
		cfg: cfg,
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		for range tickerRun.C {
			log.Info("deposit")

		}
		return nil
	})
	return nil
}
