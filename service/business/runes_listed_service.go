package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
)

type RunesListedService interface {
	RunesListedList(orderByType, orderBy, tokenAddress string, pageNum, pageSize int) ([]business.RunesListed, int64)
	RunesListedInfo(orderId string) business.RunesListed
}

type runesListedService struct {
	db *database.DB
}

func NewRunesListedService(db *database.DB) RunesListedService {
	return &runesListedService{db: db}
}

func (r runesListedService) RunesListedList(orderByType, orderBy, tokenAddress string, pageNum, pageSize int) ([]business.RunesListed, int64) {
	return r.db.RunesListed.ListRunesListed(orderByType, orderBy, tokenAddress, pageNum, pageSize)
}

func (r runesListedService) RunesListedInfo(orderId string) business.RunesListed {
	return r.db.RunesListed.RunesListedInfo(orderId)
}
