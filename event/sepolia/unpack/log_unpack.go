package unpack

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/global_const"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
	common2 "github.com/Pump-Elf-Ranch/per_apollo/database/common"
	abi "github.com/Pump-Elf-Ranch/per_apollo/event/sepolia/abi"
)

var (
	PerTokenBase, _ = abi.NewPerTokenBaseFilterer(common.Address{}, nil)
)

func ItemMinted(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := PerTokenBase.ParseItemMinted(*rlpLog)
	if unpackErr != nil {
		log.Error("parse list error", "error", unpackErr)
		return unpackErr
	}
	to := uEvent.To
	mintListed := business.MintListed{
		MintAddress:     to,
		Timestamp:       event.Timestamp,
		Nonce:           uEvent.Nonce,
		ContractAddress: event.ContractAddress,
		MintType:        uEvent.MintType,
		BlockNumber:     event.BlockNumber,
		TxHash:          rlpLog.TxHash,
	}
	storeErr := db.MintListed.StoreMintListed(mintListed)
	if storeErr != nil {
		log.Error("store mint listed error", "error", storeErr)
		return storeErr
	}
	return nil
}

func ItemBought(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := PerTokenBase.ParseItemBought(*rlpLog)
	if unpackErr != nil {
		log.Error("parse list error", "error", unpackErr)
		return unpackErr
	}
	buyPropListed := business.BuyPropListed{
		Buyer:           uEvent.Buyer,
		ItemId:          uEvent.ItemId,
		Timestamp:       event.Timestamp,
		Price:           uEvent.Price,
		ItemType:        uEvent.ItemType,
		Pid1:            uEvent.Pid1,
		Pid2:            uEvent.Pid2,
		RanchId:         uEvent.RanchId,
		IsDeposit:       global_const.NoDeposit,
		ContractAddress: event.ContractAddress,
		BlockNumber:     event.BlockNumber,
		TxHash:          rlpLog.TxHash,
	}

	storeErr := db.BuyPropListed.StoreBuyProp(buyPropListed)
	if storeErr != nil {
		log.Error("store buy prop listed error", "error", storeErr)
		return storeErr
	}
	return nil
}
