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
	ItemType        string         `json:"itemType"`
	IsDeposit       int            `json:"isDeposit"`
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
}

type BuyPropListedDB interface {
	BuyPropListedView
	StoreBuyProp(BuyPropListed) error
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
