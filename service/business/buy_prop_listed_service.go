package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
)

type BuyPropService interface {
	ListBuyPropListed(userAddress string, pageNum, pageSize int) ([]business.BuyPropListed, int64)
}

type buyPropService struct {
	db *database.DB
}

func NewBuyPropService(db *database.DB) BuyPropService {
	return &buyPropService{db: db}
}

func (r buyPropService) ListBuyPropListed(userAddress string, pageNum, pageSize int) ([]business.BuyPropListed, int64) {
	return r.db.BuyPropListed.ListBuyPropListed(userAddress, pageNum, pageSize)
}
