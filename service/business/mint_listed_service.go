package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
)

type MintListedService interface {
	MintListedList(userAddress string, pageNum, pageSize int) ([]business.MintListed, int64)
}

type mintListedService struct {
	db *database.DB
}

func NewMintListedService(db *database.DB) MintListedService {
	return &mintListedService{db: db}
}

func (r mintListedService) MintListedList(userAddress string, pageNum, pageSize int) ([]business.MintListed, int64) {
	return r.db.MintListed.ListMintListed(userAddress, pageNum, pageSize)
}
