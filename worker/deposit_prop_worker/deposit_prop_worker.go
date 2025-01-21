package deposit_prop_worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/DelphinusLab/zkwasm-minirollup-rpc-go/zkwasm"
	"github.com/Pump-Elf-Ranch/per_apollo/common/global_const"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/tasks"
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
)

type WorkerProcessor struct {
	db        *database.DB
	tasks     tasks.Group
	cfg       *config.Config
	zkwasmRpc *zkwasm.ZKWasmAppRpc
}

var DEPOSIT_CMD = big.NewInt(8)
var INIT_CMD = big.NewInt(1)

func NewWorkerProcessor(db *database.DB, shutdown context.CancelCauseFunc, cfg *config.Config,
	zkwasmRpc *zkwasm.ZKWasmAppRpc) (*WorkerProcessor, error) {
	workerProcessor := WorkerProcessor{
		db: db,
		tasks: tasks.Group{HandleCrit: func(err error) {
			shutdown(fmt.Errorf("critical error in worker processor: %w", err))
		}},
		cfg:       cfg,
		zkwasmRpc: zkwasmRpc,
	}
	return &workerProcessor, nil
}

func (b *WorkerProcessor) WorkerStart() error {
	tickerRun := time.NewTicker(time.Second * 5)
	b.tasks.Go(func() error {
		var adminExist = false
		for range tickerRun.C {
			log.Info("deposit")
			adminExist = b.AdminExist()
			if adminExist {
				b.DepositProp()
			} else {
				b.InitAdmin()
			}
		}
		return nil
	})
	return nil
}

func (b *WorkerProcessor) DepositProp() {
	adminKey := b.cfg.AdminKey
	buyPropListed := b.db.BuyPropListed.ListNeedDeposit()
	for _, buyProp := range buyPropListed {
		nonce, getNonceErr := b.zkwasmRpc.GetNonce(adminKey)
		if getNonceErr != nil {
			log.Error("GetNonce", "error", getNonceErr)
			buyProp.TryTimes++
			err := b.db.BuyPropListed.UpdateBuyProp(buyProp)
			if err != nil {
				log.Error("UpdateBuyProp", "error", err)
			}
			continue
		}
		depositCmd := b.zkwasmRpc.CreateCommand(nonce, DEPOSIT_CMD, big.NewInt(0))
		ranchId := buyProp.RanchId
		propType := buyProp.ItemId
		depositP := new(big.Int).Lsh(ranchId, 32)
		depositP.Add(depositP, propType)
		cmd := [4]*big.Int{
			depositCmd,
			buyProp.Pid1,
			buyProp.Pid2,
			depositP,
		}
		resultDeposit, err := b.zkwasmRpc.SendTransaction(cmd, adminKey)
		if err != nil {
			log.Error("SendTransaction", "error", err)
			buyProp.TryTimes++
			err = b.db.BuyPropListed.UpdateBuyProp(buyProp)
			if err != nil {
				log.Error("UpdateBuyProp", "error", err)
			}
			continue
		}
		log.Info("deposit", "resultDeposit", resultDeposit)
		buyProp.TryTimes++
		buyProp.IsDeposit = global_const.IsDeposit
		err = b.db.BuyPropListed.UpdateBuyProp(buyProp)
		if err != nil {
			log.Error("UpdateBuyProp", "error", err)
		}
	}
}

func (b *WorkerProcessor) InitAdmin() {
	adminKey := b.cfg.AdminKey
	initPlayerCmd := b.zkwasmRpc.CreateCommand(big.NewInt(0), INIT_CMD, big.NewInt(0))
	cmd := [4]*big.Int{
		initPlayerCmd,
		big.NewInt(0),
		big.NewInt(0),
		big.NewInt(0),
	}
	resultInitPlayer, err := b.zkwasmRpc.SendTransaction(cmd, adminKey)
	if err != nil {
		log.Error("SendTransaction", "error", err)
		return
	}
	log.Info("initPlayer", "resultInitPlayer", resultInitPlayer)
}

func (b *WorkerProcessor) AdminExist() bool {
	state, getStateErr := b.zkwasmRpc.QueryState(b.cfg.AdminKey)
	if getStateErr != nil {
		log.Error("QueryState", "error", getStateErr)
		return false
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(state["data"].(string)), &data)
	if err != nil {
		log.Error("Unmarshal", "error", err)
		return false
	} else {
		player, ok := data["player"]
		if ok {
			log.Info("adminExist", "admin player", player)
			return true
		}
	}
	return false
}
