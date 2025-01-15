package business

import (
	"errors"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type MintListed struct {
	GUID            uuid.UUID      `gorm:"primaryKey;DEFAULT:replace(uuid_generate_v4()::text,'-','');serializer:uuid" json:"guid"`
	MintAddress     common.Address `gorm:"serializer:bytes" json:"mintAddress"`
	Timestamp       uint64         ` json:"timestamp"`
	Nonce           *big.Int       `gorm:"serializer:u256" json:"nonce"`
	ContractAddress common.Address `gorm:"serializer:bytes" json:"contractAddress"`
	MintType        string         `json:"mintType"`
	BlockNumber     *big.Int       `gorm:"serializer:u256" json:"blockNumber"`
	TxHash          common.Hash    `gorm:"serializer:bytes" json:"txHash"`
}

func (MintListed) TableName() string {
	return "mint_nft_listed"
}

type mintListedDB struct {
	gorm *gorm.DB
}

func NewMintListedDB(db *gorm.DB) MintListedDB {
	return &mintListedDB{gorm: db}
}

type MintListedView interface {
	ListMintListed(userAddress string, pageNum, pageSize int) ([]MintListed, int64)
}

type MintListedDB interface {
	MintListedView
	StoreMintListed(MintListed) error
}

func (db mintListedDB) ListMintListed(userAddress string, pageNum, pageSize int) ([]MintListed, int64) {
	var listed []MintListed
	var count int64
	this := db.gorm.Table(MintListed{}.TableName())

	if userAddress != "" {
		this = this.Where(MintListed{MintAddress: common.HexToAddress(userAddress)})
	}
	this.Count(&count)
	if pageNum > 0 && pageSize > 0 {
		this = this.Limit(pageSize).Offset((pageNum - 1) * pageSize)
	}
	err := this.Find(&listed).Error
	if err != nil {
		log.Error("ListMintListed", "err", err.Error())
		return nil, 0
	}
	return listed, count
}

func (db mintListedDB) StoreMintListed(listed MintListed) error {
	if listed.TxHash.String() == "" {
		return nil
	}
	var exits MintListed
	err := db.gorm.Table(listed.TableName()).Where(MintListed{TxHash: listed.TxHash}).Take(&exits).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			result := db.gorm.Table(listed.TableName()).Omit("guid").Create(&listed)
			return result.Error
		}
	}
	return nil
}
