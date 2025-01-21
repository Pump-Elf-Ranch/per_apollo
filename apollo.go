package PerSynchronizer

import (
	"context"
	"errors"
	"fmt"
	"github.com/DelphinusLab/zkwasm-minirollup-rpc-go/zkwasm"
	"github.com/Pump-Elf-Ranch/per_apollo/event/sepolia"
	"github.com/Pump-Elf-Ranch/per_apollo/worker/clean_data_worker"
	"github.com/Pump-Elf-Ranch/per_apollo/worker/deposit_prop_worker"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"

	business_api "github.com/Pump-Elf-Ranch/per_apollo/api/business"
	"github.com/Pump-Elf-Ranch/per_apollo/common/errors_h"
	"github.com/Pump-Elf-Ranch/per_apollo/common/global_const"
	"github.com/Pump-Elf-Ranch/per_apollo/common/middleware"
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/service"
	"github.com/Pump-Elf-Ranch/per_apollo/synchronizer"
	"github.com/Pump-Elf-Ranch/per_apollo/synchronizer/node"
)

type PerApollo struct {
	shutdown                context.CancelCauseFunc
	db                      *database.DB
	Synchronizer            map[uint64]*synchronizer.Synchronizer
	stopped                 atomic.Bool
	chainIdList             []uint64
	ethClient               map[uint64]node.EthClient
	sepoliaEventProcessor   *sepolia.SepoliaEventProcessor
	cleanBlockHerdersWorker *clean_data_worker.WorkerProcessor
	depositPropWorker       *deposit_prop_worker.WorkerProcessor
}

func NewApi(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*PerApollo, error) {
	log.Info("NewRunesApi start️ 🕖")
	out := &PerApollo{
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.newApi(cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	log.Info("PerPerApi success🏅️")
	return out, nil
}

func (e *PerApollo) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := e.initDB(cfg); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	return nil
}

func (e *PerApollo) initEvent(ctx context.Context, cfg *config.Config) error {
	for i := range cfg.RPCs {
		rpc := cfg.RPCs[i]
		chainId := rpc.ChainId
		chainIdStr := fmt.Sprintf("%d", chainId)
		log.Info("Init event processor", "chainId", chainId)
		if chainId == global_const.SepoliaTestNetChainId {
			var epoch uint64 = 10_000
			var loopInterval time.Duration = time.Second * 5
			sepoliaEventProcessor, err := sepolia.NewEventProcessor(e.db,
				loopInterval,
				rpc.Contracts, chainIdStr, rpc.StartBlock, rpc.EventStartBlock, epoch, e.shutdown, rpc.RpcUrl)
			if err != nil {
				return fmt.Errorf("failed to init sepolia event processor: %w", err)
			}
			e.sepoliaEventProcessor = sepoliaEventProcessor
		}
	}
	return nil
}

func (e *PerApollo) initRPCClients(ctx context.Context, conf *config.Config) error {
	for i := range conf.RPCs {
		log.Info("Init rpc client", "ChainId", conf.RPCs[i].ChainId, "RpcUrl", conf.RPCs[i].RpcUrl)
		rpc := conf.RPCs[i]
		ethClient, err := node.DialEthClient(ctx, rpc.RpcUrl)
		if err != nil {
			log.Error("dial eth client fail", "err", err)
			return fmt.Errorf("Failed to dial L1 client: %w", err)
		}
		if e.ethClient == nil {
			e.ethClient = make(map[uint64]node.EthClient)
		}
		e.ethClient[rpc.ChainId] = ethClient
		e.chainIdList = append(e.chainIdList, rpc.ChainId)
	}
	log.Info("Init rpc client success")
	return nil
}

func (e *PerApollo) initDB(cfg *config.Config) error {
	db, err := database.NewDB(cfg.MasterDb)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	e.db = db
	log.Info("Init database success")
	return nil
}

func NewSync(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*PerApollo, error) {
	log.Info("NewSync start 🕖")
	out := &PerApollo{
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.initRPCClients(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.initEvent(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.initCleanBlockWorker(cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.initDepositWorker(cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	if err := out.newSync(cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	log.Info("NewSync success 🏅️")
	return out, nil
}

func (e *PerApollo) initCleanBlockWorker(cfg *config.Config) error {
	cleanDataWorker, err := clean_data_worker.NewWorkerProcessor(e.db, e.shutdown, cfg)
	if err != nil {
		return fmt.Errorf("failed to init  to clean_data_woker: %w", err)
	}
	e.cleanBlockHerdersWorker = cleanDataWorker
	log.Info("Init clean_data_woker success")
	return nil
}

func (e *PerApollo) initDepositWorker(cfg *config.Config) error {
	zkwamRpc := zkwasm.NewZKWasmAppRpc(cfg.ZkwasmRpcUrl)
	depositPropWorker, err := deposit_prop_worker.NewWorkerProcessor(e.db, e.shutdown, cfg, zkwamRpc)
	if err != nil {
		return fmt.Errorf("failed to init  to depositPropWorker: %w", err)
	}
	e.depositPropWorker = depositPropWorker
	log.Info("Init depositPropWorker success")
	return nil
}

func (e *PerApollo) newApi(cfg *config.Config) error {
	// init base service
	service.NewBaseService(e.db, cfg)
	gin.ForceConsoleColor()
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(errors_h.Recover)
	r.Use(middleware.Cors())
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "health",
		})
	})

	business_api.BuyPropApi(r)
	business_api.MintListedApi(r)
	business_api.ContractInfoApi(r)

	port := fmt.Sprintf(":%d", cfg.Server.Port)
	err := r.Run(port)
	if err != nil {
		log.Error("failed to start api server", "err", err)
		return err
	}
	return nil
}
func (e *PerApollo) newSync(cfg *config.Config) error {
	for i := range cfg.RPCs {
		log.Info("Init synchronizer success", "chainId", cfg.RPCs[i].ChainId)
		rpcItem := cfg.RPCs[i]
		thisConfig := synchronizer.Config{
			LoopIntervalMsec:  5,
			HeaderBufferSize:  uint(rpcItem.HeaderBufferSize),
			ConfirmationDepth: big.NewInt(int64(1)),
			StartHeight:       big.NewInt(int64(rpcItem.StartBlock)),
			ChainId:           uint(rpcItem.ChainId),
		}
		synchronizerTemp, err := synchronizer.NewSynchronizer(&thisConfig, e.db, rpcItem.Contracts, e.ethClient[rpcItem.ChainId], e.shutdown)
		if err != nil {
			return err
		}
		if e.Synchronizer == nil {
			e.Synchronizer = make(map[uint64]*synchronizer.Synchronizer)
		}
		e.Synchronizer[rpcItem.ChainId] = synchronizerTemp
	}
	return nil
}

func (e *PerApollo) Start(ctx context.Context) error {
	for i := range e.chainIdList {
		log.Info("starting Sync", "chainId", e.chainIdList[i])
		realChainId := e.chainIdList[i]
		if err := e.Synchronizer[realChainId].Start(); err != nil {
			return fmt.Errorf("failed to start L1 Sync: %w", err)
		}
	}
	err := e.sepoliaEventProcessor.Start()
	if err != nil {
		return fmt.Errorf("failed to start sepolia event processor: %w", err)
	}
	err = e.cleanBlockHerdersWorker.WorkerStart()
	if err != nil {
		return fmt.Errorf("failed to start clean block herders worker: %w", err)
	}
	err = e.depositPropWorker.WorkerStart()
	if err != nil {
		return fmt.Errorf("failed to start deposit prop worker: %w", err)
	}
	return nil
}

func (e *PerApollo) initConfig(ctx context.Context, cfg *config.Config) error {
	db, err := database.NewDB(cfg.MasterDb)
	if err != nil {
		log.Error("failed to connect to database", "err", err)
		return err
	}
	e.db = db
	return nil
}

func (e *PerApollo) Stop(ctx context.Context) error {
	var result error

	for i := range e.chainIdList {
		if e.Synchronizer[e.chainIdList[i]] != nil {
			if err := e.Synchronizer[e.chainIdList[i]].Close(); err != nil {
				result = errors.Join(result, fmt.Errorf("failed to close L2 Synchronizer: %w", err))
			}
		}
		if e.ethClient[e.chainIdList[i]] != nil {
			e.ethClient[e.chainIdList[i]].Close()
		}
		if e.sepoliaEventProcessor != nil {
			err := e.sepoliaEventProcessor.Close()
			if err != nil {
				result = errors.Join(result, fmt.Errorf("failed to close sepolia event processor: %w", err))
			}
		}
	}

	log.Info("apollo stopped")
	e.stopped.Store(true)
	return nil
}

func (e *PerApollo) Stopped() bool {
	return e.stopped.Load()
}
