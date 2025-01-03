package business

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type RunesActivity struct {
	GUID                  uuid.UUID      `gorm:"primaryKey;DEFAULT replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	Ticker                string         `json:"ticker"`
	Symbol                string         `json:"symbol"`
	TokenAddress          common.Address `gorm:"serializer:bytes" json:"tokenAddress"`
	TokenDecimal          int            `json:"tokenDecimal"`
	Amount                string         ` json:"amount"`
	Price                 string         ` json:"price"`
	Total                 string         ` json:"total"`
	Seller                common.Address `gorm:"serializer:bytes" json:"seller"`
	Buyer                 common.Address `gorm:"serializer:bytes" json:"buyer"`
	OrderId               *big.Int       `gorm:"serializer:u256" json:"orderId"`
	ActivityType          int            `json:"activity_type"`
	Timestamp             uint64         `json:"timestamp"`
	MarketContractAddress common.Address `gorm:"serializer:bytes" json:"marketContractAddress"`
	BlockNumber           *big.Int       `gorm:"serializer:u256" json:"blockNumber"`
	TxHash                common.Hash    `gorm:"serializer:bytes" json:"txHash"`
}

func (RunesActivity) TableName() string {
	return "runes_activity"
}

type runesActivityDB struct {
	gorm *gorm.DB
}

func NewRunesActivityDB(db *gorm.DB) RunesActivityDB {
	return &runesActivityDB{gorm: db}
}

type RunesActivityView interface {
	ListRunesActivity(activityType []int, pageNum, pageSize int) ([]RunesActivity, int64)
}

type RunesActivityDB interface {
	RunesActivityView
	StoreRunesActivity(RunesActivity) error
	UpdateRunesActivity(RunesActivity) error
}

func (db runesActivityDB) ListRunesActivity(activityType []int, pageNum, pageSize int) ([]RunesActivity, int64) {
	var activities []RunesActivity
	var count int64
	this := db.gorm.Table(RunesActivity{}.TableName())
	if len(activityType) > 0 {
		this = this.Where("activity_type in ?", activityType)
	}
	this.Count(&count)
	result := this.Order("timestamp desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&activities)
	if result.Error != nil {
		log.Error("ListRunesActivity", "error", result.Error)

	}
	return activities, count
}

func (db runesActivityDB) StoreRunesActivity(activity RunesActivity) error {
	if activity.TxHash.String() == "" {
		return nil
	}
	var exits RunesActivity
	err := db.gorm.Table(activity.TableName()).Where(RunesActivity{TxHash: activity.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(activity.TableName()).Omit("guid").Create(&activity)
			return result.Error
		}
	}
	return nil
}

func (db runesActivityDB) UpdateRunesActivity(activity RunesActivity) error {
	//TODO implement me
	panic("implement me")
}
