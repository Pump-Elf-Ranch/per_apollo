package business

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type RunesOrder struct {
	GUID                  uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	Ticker                string         `json:"ticker"`
	Symbol                string         `json:"symbol"`
	TokenDecimal          int            `json:"tokenDecimal"`
	TokenAddress          common.Address `gorm:"serializer:bytes" json:"token_address"`
	Amount                string         ` json:"amount"`
	Price                 string         ` json:"price"`
	Total                 string         ` json:"total"`
	Seller                common.Address `gorm:"serializer:bytes" json:"seller"`
	Buyer                 common.Address `gorm:"serializer:bytes" json:"buyer"`
	OrderId               *big.Int       `gorm:"serializer:u256" json:"orderId"`
	Status                int            `json:"status"`
	OrderType             int8           `json:"orderType"`
	Timestamp             uint64         `json:"timestamp"`
	MarketContractAddress common.Address `gorm:"serializer:bytes" json:"market_contract_address"`
	BlockNumber           *big.Int       `gorm:"serializer:u256" json:"blockNumber"`
	TxHash                common.Hash    `gorm:"serializer:bytes" json:"txHash"`
	OrderTxHash           common.Hash    `gorm:"serializer:bytes" json:"orderTxHash"`
}

func (RunesOrder) TableName() string {
	return "runes_order"
}

type runesOrderDB struct {
	gorm *gorm.DB
}

func NewRunesOrderDB(db *gorm.DB) RunesOrderDB {
	return &runesOrderDB{gorm: db}
}

type RunesOrderView interface {
	ListRunesOrder(userAddress, tokenAddress,
		orderId string, pageNum, pageSize int) ([]RunesOrder, int64)
	RunesOrderInfo(orderId string) RunesOrder
}

type RunesOrderDB interface {
	RunesOrderView
	StoreRunesOrder(RunesOrder) error
	UpdateRunesOrderAmountPriceByOrderId(RunesOrder) error
	UpdateRunesOrderStatus(orderId *big.Int, status int,
		buyer common.Address, orderTxHash common.Hash) error
}

func (db runesOrderDB) ListRunesOrder(userAddress, tokenAddress,
	orderId string, pageNum, pageSize int) ([]RunesOrder, int64) {
	var orders []RunesOrder
	var count int64
	this := db.gorm.Table(RunesOrder{}.TableName())
	if userAddress != "" {
		this = this.Where(RunesOrder{Seller: common.HexToAddress(userAddress)}).
			Or(RunesOrder{Buyer: common.HexToAddress(userAddress)})
	}
	if tokenAddress != "" {
		this = this.Where(RunesOrder{TokenAddress: common.HexToAddress(tokenAddress)})
	}
	if orderId != "" {
		this = this.Where("order_id = ?", orderId)
	}
	this = this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	this = this.Find(&orders)
	if this.Error != nil {
		log.Error("ListRunesOrder", "err", this.Error)
	}
	return orders, count
}

func (db runesOrderDB) StoreRunesOrder(order RunesOrder) error {
	if order.TxHash.String() == "" {
		return nil
	}
	var exits RunesOrder
	err := db.gorm.Table(order.TableName()).Where(RunesOrder{TxHash: order.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(order.TableName()).Omit("guid").Create(&order)
			return result.Error
		}
	}
	return nil

}

func (db runesOrderDB) UpdateRunesOrderStatus(orderId *big.Int, status int,
	buyer common.Address, orderTxHash common.Hash) error {

	runesOrder := RunesOrder{
		Status: status,
	}
	if buyer.String() != "" {
		runesOrder.Buyer = buyer
	}
	if orderTxHash.String() != "" {
		runesOrder.OrderTxHash = orderTxHash
	}

	result := db.gorm.Table(RunesOrder{}.TableName()).Where("order_id = ?", orderId).
		Updates(runesOrder)
	if result.Error != nil {
		log.Error("UpdateRunesOrderStatus", "err", result.Error)
		return result.Error
	}
	return nil
}

func (db runesOrderDB) RunesOrderInfo(orderId string) RunesOrder {
	var order RunesOrder
	err := db.gorm.Table(RunesOrder{}.TableName()).Where("order_id = ?", orderId).Take(&order).Error
	if err != nil {
		log.Error("RunesOrderInfo", "err", err)
	}
	return order
}

func (db runesOrderDB) UpdateRunesOrderAmountPriceByOrderId(runesOrder RunesOrder) error {
	result := db.gorm.Table(RunesOrder{}.TableName()).Where(RunesOrder{OrderId: runesOrder.OrderId,
		TokenAddress: runesOrder.TokenAddress}).Updates(RunesOrder{Amount: runesOrder.Amount, Price: runesOrder.Price,
		Total: runesOrder.Total})
	if result.Error != nil {
		log.Error("UpdateRunesOrderAmountPriceByOrderId", "err", result.Error)
		return result.Error
	}
	return nil
}
