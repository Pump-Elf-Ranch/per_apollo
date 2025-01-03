package bitlayer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/bigint"
	"github.com/Pump-Elf-Ranch/per_apollo/common/tasks"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/block_listener"
	common2 "github.com/Pump-Elf-Ranch/per_apollo/database/common"
	abi "github.com/Pump-Elf-Ranch/per_apollo/event/bitlayer/abi"
	"github.com/Pump-Elf-Ranch/per_apollo/event/bitlayer/unpack"
)

type BitLayerEventProcessor struct {
	loopInterval    time.Duration
	resourceCtx     context.Context
	resourceCancel  context.CancelFunc
	tasks           tasks.Group
	db              *database.DB
	epoch           uint64
	eventStartBlock uint64
	startHeight     *big.Int
	contracts       []string
	chainId         string
	ethClient       *ethclient.Client
}

func NewEventProcessor(db *database.DB, loopInterval time.Duration, contracts []string,
	chainId string,
	startHeight uint64, eventStartBlock uint64, epoch uint64,
	shutdown context.CancelCauseFunc,
	rpcUrl string) (*BitLayerEventProcessor, error) {
	ethClient, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return nil, err
	}
	resCtx, resCancel := context.WithCancel(context.Background())
	return &BitLayerEventProcessor{
		db:             db,
		resourceCtx:    resCtx,
		resourceCancel: resCancel,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error processor: %w", err))
		}},
		startHeight:     new(big.Int).SetUint64(startHeight),
		eventStartBlock: eventStartBlock,
		chainId:         chainId,
		contracts:       contracts,
		loopInterval:    loopInterval,
		epoch:           epoch,
		ethClient:       ethClient,
	}, nil
}

func (bep *BitLayerEventProcessor) Start() error {
	tickerEventOn1 := time.NewTicker(bep.loopInterval)
	log.Info("starting bitlayer event processor...")
	bep.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := bep.onData()
			if err != nil {
				log.Error("l2 event processor error", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (bep *BitLayerEventProcessor) onData() error {
	if bep.startHeight == nil {
		lastListenBlock, err := bep.db.BlockListener.GetLastBlockNumber(bep.chainId)
		if err != nil {
			log.Error("failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			lastListenBlock = &block_listener.BlockListener{
				ChainId:     bep.chainId,
				BlockNumber: big.NewInt(0),
			}
		}
		bep.startHeight = lastListenBlock.BlockNumber
		if bep.startHeight.Cmp(big.NewInt(int64(bep.eventStartBlock))) == -1 {
			bep.startHeight = big.NewInt(int64(bep.eventStartBlock))
		}
	} else {
		bep.startHeight = new(big.Int).Add(bep.startHeight, bigint.One)
	}
	fromHeight := bep.startHeight
	toHeight := new(big.Int).Add(fromHeight, big.NewInt(int64(bep.epoch)))

	latestBlockHeader, err := bep.db.Blocks.ChainLatestBlockHeader(bep.chainId)
	if err != nil {
		bep.startHeight = new(big.Int).Sub(bep.startHeight, bigint.One)
		return err
	}
	if latestBlockHeader == nil {
		bep.startHeight = new(big.Int).Sub(bep.startHeight, bigint.One)
		return nil
	}
	if latestBlockHeader.Number.Cmp(fromHeight) == -1 {
		fromHeight = latestBlockHeader.Number
		toHeight = latestBlockHeader.Number
	} else {
		if latestBlockHeader.Number.Cmp(toHeight) == -1 {
			toHeight = latestBlockHeader.Number
		}
	}
	if err := bep.db.Transaction(func(tx *database.DB) error {
		eventsFetchErr := bep.eventsFetch(fromHeight, toHeight)
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := block_listener.BlockListener{
			ChainId:     bep.chainId,
			BlockNumber: toHeight,
		}
		updateErr := bep.db.BlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Error("update last block", " err :", updateErr)
			return updateErr
		}
		return nil
	}); err != nil {
		bep.startHeight = new(big.Int).Sub(bep.startHeight, bigint.One)
		return err
	}
	bep.startHeight = toHeight
	return nil
}

func (bep *BitLayerEventProcessor) eventsFetch(fromHeight, toHeight *big.Int) error {
	contracts := bep.contracts
	for _, contract := range contracts {
		contractEventFilter := common2.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		events, err := bep.db.ContractEvents.ContractEventsWithFilter(bep.chainId, contractEventFilter, fromHeight, toHeight)
		if err != nil {
			log.Error("failed to index ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := bep.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (bep *BitLayerEventProcessor) eventUnpack(event common2.ContractEvent) error {
	tokenMarketAbi, _ := abi.L2TokenMarketMetaData.GetAbi()
	switch event.EventSignature.String() {
	case tokenMarketAbi.Events["List"].ID.String():
		err := unpack.Listed(event, bep.db, bep.ethClient)
		return err
	case tokenMarketAbi.Events["UnList"].ID.String():
		err := unpack.UnList(event, bep.db, bep.ethClient)
		return err
	case tokenMarketAbi.Events["OrderConfirm"].ID.String():
		err := unpack.OrderConfirm(event, bep.db, bep.ethClient)
		return err
	case tokenMarketAbi.Events["Update"].ID.String():
		err := unpack.Update(event, bep.db, bep.ethClient)
		return err
	}
	return nil
}

func (bep *BitLayerEventProcessor) Close() error {
	bep.resourceCancel()
	return bep.tasks.Wait()
}
