package unpack

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/database"
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
	fmt.Println(uEvent)

	return nil
}

func ItemBought(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	return nil
}
