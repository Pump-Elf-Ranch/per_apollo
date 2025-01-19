package clean_data_worker

import (
	"context"
	"fmt"
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/tasks"
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
			log.Info("clean table ")
			for i := 0; i < len(b.cfg.RPCs); i++ {
				rpc := b.cfg.RPCs[i]
				chainId := rpc.ChainId
				log.Info("clean table by chainId", "chainId", chainId)
				err := b.db.Blocks.CleanBlockHerders(strconv.Itoa(int(chainId)))
				if err != nil {
					log.Error("clean table by chainId", "chainId", chainId, "err ", err)
					continue
				}
			}

		}
		return nil
	})
	return nil
}
