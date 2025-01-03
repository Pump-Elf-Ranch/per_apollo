package business

import (
	"errors"
	"github.com/Pump-Elf-Ranch/per_apollo/common/global_const"
	"time"

	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type RunesListed struct {
	GUID                  uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	Ticker                string         `json:"ticker"`
	Symbol                string         `json:"symbol"`
	TokenAddress          common.Address `gorm:"serializer:bytes" json:"tokenAddress"`
	TokenDecimal          int            `json:"tokenDecimal"`
	Amount                string         ` json:"amount"`
	Price                 string         ` json:"price"`
	Total                 string         ` json:"total"`
	Seller                common.Address `gorm:"serializer:bytes" json:"seller"`
	OrderId               *big.Int       `gorm:"serializer:u256" json:"orderId"`
	Status                int            `json:"status"`
	CreateTime            uint64         `json:"create_time"`
	UpdateTime            uint64         `json:"update_time"`
	MarketContractAddress common.Address `gorm:"serializer:bytes" json:"marketContractAddress"`
	BlockNumber           *big.Int       `gorm:"serializer:u256" json:"blockNumber"`
	TxHash                common.Hash    `gorm:"serializer:bytes" json:"txHash"`
}

func (RunesListed) TableName() string {
	return "runes_listed"
}

type runesListedDB struct {
	gorm *gorm.DB
}

func NewRunesListedDB(db *gorm.DB) RunesListedDB {
	return &runesListedDB{gorm: db}
}

type RunesListedView interface {
	ListRunesListed(orderByType, orderBy, tokenAddress string, pageNum, pageSize int) ([]RunesListed, int64)
	RunesListedInfo(orderId string) RunesListed
}

type RunesListedDB interface {
	RunesListedView
	StoreRunesListed(RunesListed) error
	UpdateRunesListed(RunesListed) error
	UpdateRunesListedAmountPriceByOrderId(RunesListed) error
	UpdateRunesListedStatus(orderId *big.Int, status int) error
}

func (db runesListedDB) ListRunesListed(orderByType, orderBy, tokenAddress string, pageNum, pageSize int) ([]RunesListed, int64) {
	var listed []RunesListed
	var count int64
	this := db.gorm.Table(RunesListed{}.TableName())
	if orderByType != "" {
		if orderBy == "asc" {
			this = this.Order(orderByType + " asc")
		} else {
			this = this.Order(orderByType + " desc")
		}
	} else {
		this = this.Order("create_time desc")
	}
	if tokenAddress != "" {
		this = this.Where(RunesListed{TokenAddress: common.HexToAddress(tokenAddress)})
	}
	this.Where("status = ?", global_const.Selling)
	this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	err := this.Find(&listed).Error
	if err != nil {
		log.Error("ListRunesListed", "err", err.Error())
		return nil, 0
	}
	return listed, count
}

func (db runesListedDB) StoreRunesListed(listed RunesListed) error {
	if listed.TxHash.String() == "" {
		return nil
	}
	var exits RunesListed
	err := db.gorm.Table(listed.TableName()).Where(RunesListed{TxHash: listed.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(listed.TableName()).Omit("guid").Create(&listed)
			return result.Error
		}
	}
	return nil
}

func (db runesListedDB) UpdateRunesListed(listed RunesListed) error {
	//TODO implement me
	panic("implement me")
}

func (db runesListedDB) RunesListedInfo(orderId string) RunesListed {
	var listed RunesListed
	result := db.gorm.Table(RunesListed{}.TableName()).Where("order_id = ?", orderId).Take(&listed)
	if result.Error != nil {
		log.Error("RunesListedInfo", "err", result.Error)
	}
	return listed
}

func (db runesListedDB) UpdateRunesListedStatus(orderId *big.Int, status int) error {
	result := db.gorm.Table(RunesListed{}.TableName()).Where("order_id = ?", orderId).Update("status", status).
		Update("update_time", uint64(time.Now().Unix()))
	if result.Error != nil {
		log.Error("UpdateRunesListedStatus", "err", result.Error)
		return result.Error
	}
	return nil
}

func (db runesListedDB) UpdateRunesListedAmountPriceByOrderId(runesListed RunesListed) error {
	result := db.gorm.Table(RunesListed{}.TableName()).Where(RunesListed{OrderId: runesListed.OrderId,
		TokenAddress: runesListed.TokenAddress}).Updates(RunesListed{Amount: runesListed.Amount, Price: runesListed.Price,
		Total: runesListed.Total, UpdateTime: runesListed.UpdateTime})
	if result.Error != nil {
		log.Error("UpdateRunesListedAmountPriceByOrderId", "err", result.Error)
		return result.Error
	}
	return nil
}
