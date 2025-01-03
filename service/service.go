package service

import (
	"github.com/Pump-Elf-Ranch/per_apollo/config"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	businessService "github.com/Pump-Elf-Ranch/per_apollo/service/business"
	"github.com/ethereum/go-ethereum/log"
)

type service struct {
	Db                   *database.DB
	Cfg                  *config.Config
	RunesActivityService businessService.RunesActivityService
	RunesListedService   businessService.RunesListedService
	RunesOrderService    businessService.RunesOrderService
}

var BaseService *service

func NewBaseService(db *database.DB, cfg *config.Config) {
	BaseService = &service{
		Db:                   db,
		Cfg:                  cfg,
		RunesActivityService: businessService.NewRunesActivityService(db),
		RunesListedService:   businessService.NewRunesListedService(db),
		RunesOrderService:    businessService.NewRunesOrderService(db),
	}
	log.Info("init BaseService success🏅️")
}
