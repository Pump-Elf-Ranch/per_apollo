package sepolia

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
	abi "github.com/Pump-Elf-Ranch/per_apollo/event/sepolia/abi"
	"github.com/Pump-Elf-Ranch/per_apollo/event/sepolia/unpack"
)

type SepoliaEventProcessor struct {
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
	rpcUrl string) (*SepoliaEventProcessor, error) {
	ethClient, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		return nil, err
	}
	resCtx, resCancel := context.WithCancel(context.Background())
	return &SepoliaEventProcessor{
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

func (sep *SepoliaEventProcessor) Start() error {
	tickerEventOn1 := time.NewTicker(sep.loopInterval)
	log.Info("starting sepolia event processor...")
	sep.tasks.Go(func() error {
		for range tickerEventOn1.C {
			err := sep.onData()
			if err != nil {
				log.Error("l2 event processor error", "err", err)
				continue
			}
		}
		return nil
	})
	return nil
}

func (sep *SepoliaEventProcessor) onData() error {
	if sep.startHeight == nil {
		lastListenBlock, err := sep.db.BlockListener.GetLastBlockNumber(sep.chainId)
		if err != nil {
			log.Error("failed to get last block heard", "err", err)
			return err
		}
		if lastListenBlock == nil {
			lastListenBlock = &block_listener.BlockListener{
				ChainId:     sep.chainId,
				BlockNumber: big.NewInt(0),
			}
		}
		sep.startHeight = lastListenBlock.BlockNumber
		if sep.startHeight.Cmp(big.NewInt(int64(sep.eventStartBlock))) == -1 {
			sep.startHeight = big.NewInt(int64(sep.eventStartBlock))
		}
	} else {
		sep.startHeight = new(big.Int).Add(sep.startHeight, bigint.One)
	}
	fromHeight := sep.startHeight
	toHeight := new(big.Int).Add(fromHeight, big.NewInt(int64(sep.epoch)))

	latestBlockHeader, err := sep.db.Blocks.ChainLatestBlockHeader(sep.chainId)
	if err != nil {
		sep.startHeight = new(big.Int).Sub(sep.startHeight, bigint.One)
		return err
	}
	if latestBlockHeader == nil {
		sep.startHeight = new(big.Int).Sub(sep.startHeight, bigint.One)
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
	if err := sep.db.Transaction(func(tx *database.DB) error {
		eventsFetchErr := sep.eventsFetch(fromHeight, toHeight)
		if eventsFetchErr != nil {
			return eventsFetchErr
		}
		lastBlock := block_listener.BlockListener{
			ChainId:     sep.chainId,
			BlockNumber: toHeight,
		}
		updateErr := sep.db.BlockListener.SaveOrUpdateLastBlockNumber(lastBlock)
		if updateErr != nil {
			log.Error("update last block", " err :", updateErr)
			return updateErr
		}
		return nil
	}); err != nil {
		sep.startHeight = new(big.Int).Sub(sep.startHeight, bigint.One)
		return err
	}
	sep.startHeight = toHeight
	return nil
}

func (sep *SepoliaEventProcessor) eventsFetch(fromHeight, toHeight *big.Int) error {
	contracts := sep.contracts
	for _, contract := range contracts {
		contractEventFilter := common2.ContractEvent{ContractAddress: common.HexToAddress(contract)}
		events, err := sep.db.ContractEvents.ContractEventsWithFilter(sep.chainId, contractEventFilter, fromHeight, toHeight)
		if err != nil {
			log.Error("failed to index ContractEventsWithFilter ", "err", err)
			return err
		}
		for _, contractEvent := range events {
			unpackErr := sep.eventUnpack(contractEvent)
			if unpackErr != nil {
				log.Error("failed to index events", "unpackErr", unpackErr)
				return unpackErr
			}
		}
	}
	return nil
}

func (sep *SepoliaEventProcessor) eventUnpack(event common2.ContractEvent) error {
	tokenMarketAbi, _ := abi.PerTokenBaseMetaData.GetAbi()
	switch event.EventSignature.String() {
	case tokenMarketAbi.Events["ItemMinted"].ID.String():
		err := unpack.ItemMinted(event, sep.db, sep.ethClient)
		return err
	case tokenMarketAbi.Events["ItemBought"].ID.String():
		err := unpack.ItemBought(event, sep.db, sep.ethClient)
		return err
	}
	return nil
}

func (sep *SepoliaEventProcessor) Close() error {
	sep.resourceCancel()
	return sep.tasks.Wait()
}
