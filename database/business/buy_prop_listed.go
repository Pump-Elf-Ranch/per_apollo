package business

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type BuyPropListed struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT:replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	Buyer           common.Address `gorm:"serializer:bytes" json:"buyer"`
	ItemId          *big.Int       `gorm:"serializer:u256" json:"itemId"`
	Timestamp       uint64         ` json:"timestamp"`
	Price           *big.Int       `gorm:"serializer:u256" json:"price"`
	Pid1            *big.Int       `gorm:"serializer:u256" json:"pid1"`
	Pid2            *big.Int       `gorm:"serializer:u256" json:"pid2"`
	RanchId         *big.Int       `gorm:"serializer:u256" json:"ranchId"`
	ItemType        string         `json:"itemType"`
	IsDeposit       int            `json:"isDeposit"`
	TryTimes        int            `json:"tryTimes"`
	ContractAddress common.Address `gorm:"serializer:bytes" json:"contractAddress"`
	BlockNumber     *big.Int       `gorm:"serializer:u256" json:"blockNumber"`
	TxHash          common.Hash    `gorm:"serializer:bytes" json:"txHash"`
}

func (BuyPropListed) TableName() string {
	return "prop_buy_listed"
}

type buyPropListedDB struct {
	gorm *gorm.DB
}

func NewBuyPropListedDB(db *gorm.DB) BuyPropListedDB {
	return &buyPropListedDB{gorm: db}
}

type BuyPropListedView interface {
	ListBuyPropListed(userAddress string, pageNum, pageSize int) ([]BuyPropListed, int64)
	ListNeedDeposit() []BuyPropListed
}

type BuyPropListedDB interface {
	BuyPropListedView
	StoreBuyProp(BuyPropListed) error
	UpdateBuyProp(buyProp BuyPropListed) error
}

func (db buyPropListedDB) ListNeedDeposit() []BuyPropListed {
	var buyPropListed []BuyPropListed
	this := db.gorm.Table(BuyPropListed{}.TableName())
	this = this.Where("is_deposit = 0 and try_times < 3")
	result := this.Order("timestamp asc").Limit(20).Find(&buyPropListed)
	if result.Error != nil {
		log.Error("ListNeedDeposit", "error", result.Error)
	}
	return buyPropListed
}

func (db buyPropListedDB) UpdateBuyProp(buyProp BuyPropListed) error {
	if buyProp.TxHash.String() == "" {
		return nil
	}
	var exits BuyPropListed
	err := db.gorm.Table(buyProp.TableName()).Where(BuyPropListed{TxHash: buyProp.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	result := db.gorm.Table(buyProp.TableName()).Where(BuyPropListed{TxHash: buyProp.TxHash}).Updates(buyProp)
	if result.Error != nil {
		log.Error("UpdateBuyProp", "error", result.Error)
	}
	return nil
}

func (db buyPropListedDB) ListBuyPropListed(userAddress string, pageNum, pageSize int) ([]BuyPropListed, int64) {
	var activities []BuyPropListed
	var count int64
	this := db.gorm.Table(BuyPropListed{}.TableName())
	if userAddress != "" {
		this = this.Where(BuyPropListed{Buyer: common.HexToAddress(userAddress)})
	}
	this.Count(&count)
	result := this.Order("timestamp desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&activities)
	if result.Error != nil {
		log.Error("ListBuyPropListed", "error", result.Error)

	}
	return activities, count
}

func (db buyPropListedDB) StoreBuyProp(buyProp BuyPropListed) error {
	if buyProp.TxHash.String() == "" {
		return nil
	}
	var exits BuyPropListed
	err := db.gorm.Table(buyProp.TableName()).Where(BuyPropListed{TxHash: buyProp.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(buyProp.TableName()).Omit("guid").Create(&buyProp)
			return result.Error
		}
	}
	return nil
}
