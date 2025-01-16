package service

import (
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	businessService "github.com/Pump-Elf-Ranch/per_apollo/service/business"
	"github.com/ethereum/go-ethereum/log"
)

type service struct {
	Db                  *database.DB
	Cfg                 *config.Config
	BuyPropService      businessService.BuyPropService
	MintListedService   businessService.MintListedService
	ContractInfoService businessService.ContractInfoService
}

var BaseService *service

func NewBaseService(db *database.DB, cfg *config.Config) {
	BaseService = &service{
		Db:                  db,
		Cfg:                 cfg,
		BuyPropService:      businessService.NewBuyPropService(db),
		MintListedService:   businessService.NewMintListedService(db),
		ContractInfoService: businessService.NewContractInfoService(cfg),
	}
	log.Info("init BaseService success🏅️")
}
