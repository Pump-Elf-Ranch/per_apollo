package business

import (
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
)

type RunesOrderService interface {
	RunesOrderList(userAddressStr, tokenAddressStr, orderId string, pageNum, pageSize int) ([]business.RunesOrder, int64)
	RunesOrderInfo(orderId string) business.RunesOrder
}

type runesOrderService struct {
	db *database.DB
}

func NewRunesOrderService(db *database.DB) RunesOrderService {
	return &runesOrderService{db: db}
}

func (r runesOrderService) RunesOrderList(userAddressStr, tokenAddressStr, orderId string, pageNum, pageSize int) ([]business.RunesOrder, int64) {
	return r.db.RunesOrder.ListRunesOrder(userAddressStr, tokenAddressStr, orderId, pageNum, pageSize)
}

func (r runesOrderService) RunesOrderInfo(orderId string) business.RunesOrder {
	return r.db.RunesOrder.RunesOrderInfo(orderId)

}
