package unpack

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/Pump-Elf-Ranch/per_apollo/common/global_const"
	"github.com/Pump-Elf-Ranch/per_apollo/database"
	"github.com/Pump-Elf-Ranch/per_apollo/database/business"
	common2 "github.com/Pump-Elf-Ranch/per_apollo/database/common"
	abi "github.com/Pump-Elf-Ranch/per_apollo/event/bitlayer/abi"
)

var (
	TokenMarketUnpack, _ = abi.NewL2TokenMarketFilterer(common.Address{}, nil)
	BtcDecimal           = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(8)), nil)
)

func Listed(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := TokenMarketUnpack.ParseList(*rlpLog)
	if unpackErr != nil {
		log.Error("parse list error", "error", unpackErr)
		return unpackErr
	}
	tokenName, tokenSymbol, tokenDecimals, unpackErr := GetTokenNameAndSymbol(uEvent.Token, ethClient)
	if unpackErr != nil {
		log.Error("get token name and symbol error", "error", unpackErr)
		return unpackErr
	}
	amount := uEvent.Amount
	price := uEvent.Price
	//tokenDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)
	//priceDecimal := new(big.Int).Div(price, BtcDecimal)
	//amountDecimal := new(big.Int).Div(amount, tokenDecimal)
	total := new(big.Int).Mul(amount, price)
	runesListed := business.RunesListed{
		Ticker:                tokenName,
		Symbol:                tokenSymbol,
		TxHash:                rlpLog.TxHash,
		TokenAddress:          uEvent.Token,
		Amount:                amount.String(),
		Price:                 price.String(),
		Total:                 total.String(),
		Seller:                uEvent.Seller,
		OrderId:               uEvent.Nonce,
		Status:                global_const.Selling,
		CreateTime:            event.Timestamp,
		UpdateTime:            event.Timestamp,
		MarketContractAddress: rlpLog.Address,
		BlockNumber:           event.BlockNumber,
		TokenDecimal:          int(tokenDecimals),
	}
	saveRunesListedErr := db.RunesListed.StoreRunesListed(runesListed)
	if saveRunesListedErr != nil {
		log.Error("store RunesListed error", "error", saveRunesListedErr)
		return saveRunesListedErr
	}

	// 保存售出订单
	runesOrder := business.RunesOrder{
		Ticker:                tokenName,
		Symbol:                tokenSymbol,
		TokenAddress:          uEvent.Token,
		Amount:                amount.String(),
		Price:                 price.String(),
		Total:                 total.String(),
		Seller:                uEvent.Seller,
		Buyer:                 common.Address{},
		OrderId:               uEvent.Nonce,
		OrderType:             global_const.Seller,
		Status:                global_const.Selling,
		Timestamp:             event.Timestamp,
		MarketContractAddress: rlpLog.Address,
		BlockNumber:           event.BlockNumber,
		TxHash:                rlpLog.TxHash,
		TokenDecimal:          int(tokenDecimals),
	}
	storeRunesErr := db.RunesOrder.StoreRunesOrder(runesOrder)
	if storeRunesErr != nil {
		log.Error("store RunesOrder error", "error", storeRunesErr)
		return storeRunesErr
	}

	return SaveActivityInfo(global_const.Listed, event, db, ethClient)
}

func UnList(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := TokenMarketUnpack.ParseUnList(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}
	// update RunesListed status
	updateRunesListedErr := db.RunesListed.UpdateRunesListedStatus(uEvent.Nonce, global_const.UnSell)
	if updateRunesListedErr != nil {
		log.Error("update RunesListed error", "error", updateRunesListedErr)
		return updateRunesListedErr
	}
	updateRunesOrderErr := db.RunesOrder.UpdateRunesOrderStatus(uEvent.Nonce, global_const.UnSell,
		common.Address{}, common.Hash{})
	if updateRunesOrderErr != nil {
		log.Error("update RunesOrder error", "error", updateRunesOrderErr)
		return updateRunesOrderErr

	}
	return SaveActivityInfo(global_const.UnListed, event, db, ethClient)
}

func OrderConfirm(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := TokenMarketUnpack.ParseOrderConfirm(*rlpLog)
	if unpackErr != nil {
		log.Error("parse order confirm error", "error", unpackErr)
		return unpackErr
	}

	updateRunesListedErr := db.RunesListed.UpdateRunesListedStatus(uEvent.Nonce, global_const.Sold)
	if updateRunesListedErr != nil {
		log.Error("update RunesListed error", "error", updateRunesListedErr)
		return updateRunesListedErr

	}
	updateRunesOrderErr := db.RunesOrder.UpdateRunesOrderStatus(uEvent.Nonce, global_const.Sold,
		uEvent.Buyer, rlpLog.TxHash)
	if updateRunesOrderErr != nil {
		log.Error("update RunesOrder error", "error", updateRunesOrderErr)
		return updateRunesOrderErr

	}
	return SaveActivityInfo(global_const.OrderConfirmed, event, db, ethClient)
}

func Update(event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	uEvent, unpackErr := TokenMarketUnpack.ParseUpdate(*rlpLog)
	if unpackErr != nil {
		return unpackErr
	}

	// update token listed
	orderId := uEvent.Nonce
	seller := uEvent.Seller
	amount := uEvent.Amount
	price := uEvent.Price
	//tokenAddress := uEvent.Token

	//_, _, tokenDecimals, unpackErr := GetTokenNameAndSymbol(tokenAddress, ethClient)
	//if unpackErr != nil {
	//	log.Error("get token name and symbol error", "error", unpackErr)
	//	return unpackErr
	//}

	//tokenDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)
	//amountDecimal := new(big.Int).Div(amount, tokenDecimal)
	//priceDecimal := new(big.Int).Div(price, BtcDecimal)
	total := new(big.Int).Mul(amount, price)

	runesListed := business.RunesListed{
		Amount:     amount.String(),
		Price:      price.String(),
		Seller:     seller,
		OrderId:    orderId,
		Total:      total.String(),
		UpdateTime: event.Timestamp,
	}
	updateErr := db.RunesListed.UpdateRunesListedAmountPriceByOrderId(runesListed)
	if updateErr != nil {
		log.Error("update RunesListed error", "error", updateErr)
		return updateErr
	}
	runesOrder := business.RunesOrder{
		Amount:  amount.String(),
		Price:   price.String(),
		Seller:  seller,
		OrderId: orderId,
		Total:   total.String(),
	}
	updateOrderErr := db.RunesOrder.UpdateRunesOrderAmountPriceByOrderId(runesOrder)
	if updateOrderErr != nil {
		log.Error("update RunesOrder error", "error", updateErr)
		return updateErr
	}
	return SaveActivityInfo(global_const.Update, event, db, ethClient)
}

func SaveActivityInfo(activityType int, event common2.ContractEvent, db *database.DB, ethClient *ethclient.Client) error {
	rlpLog := event.RLPLog
	tokenAddress := common.Address{}
	amount := big.NewInt(0)
	price := big.NewInt(0)
	seller := common.Address{}
	buyer := common.Address{}
	orderId := big.NewInt(0)
	if activityType == global_const.Listed {
		uEvent, unpackErr := TokenMarketUnpack.ParseList(*rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		tokenAddress = uEvent.Token
		amount = uEvent.Amount
		price = uEvent.Price
		seller = uEvent.Seller
		orderId = uEvent.Nonce
	} else if activityType == global_const.UnListed {
		uEvent, unpackErr := TokenMarketUnpack.ParseUnList(*rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		tokenAddress = uEvent.Token
		seller = uEvent.Seller
		orderId = uEvent.Nonce
	} else if activityType == global_const.OrderConfirmed {
		uEvent, unpackErr := TokenMarketUnpack.ParseOrderConfirm(*rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		tokenAddress = uEvent.Token
		amount = uEvent.Amount
		price = uEvent.Price
		seller = uEvent.Seller
		buyer = uEvent.Buyer
		orderId = uEvent.Nonce
	} else if activityType == global_const.Update {
		uEvent, unpackErr := TokenMarketUnpack.ParseUpdate(*rlpLog)
		if unpackErr != nil {
			return unpackErr
		}
		tokenAddress = uEvent.Token
		amount = uEvent.Amount
		price = uEvent.Price
		seller = uEvent.Seller
		orderId = uEvent.Nonce
	} else {
		return nil
	}
	tokenName, tokenSymbol, tokenDecimals, unpackErr := GetTokenNameAndSymbol(tokenAddress, ethClient)
	if unpackErr != nil {
		log.Error("get token name and symbol error", "error", unpackErr)
		return unpackErr
	}
	//tokenDecimal := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(tokenDecimals)), nil)
	//amountDecimal := new(big.Int).Div(amount, tokenDecimal)
	//priceDecimal := new(big.Int).Div(price, BtcDecimal)
	total := new(big.Int).Mul(price, amount)
	runesActivity := business.RunesActivity{
		Ticker:                tokenName,
		Symbol:                tokenSymbol,
		TokenAddress:          tokenAddress,
		Amount:                amount.String(),
		Price:                 price.String(),
		Total:                 total.String(),
		Seller:                seller,
		Buyer:                 buyer,
		OrderId:               orderId,
		ActivityType:          activityType,
		Timestamp:             event.Timestamp,
		MarketContractAddress: rlpLog.Address,
		BlockNumber:           event.BlockNumber,
		TxHash:                rlpLog.TxHash,
		TokenDecimal:          int(tokenDecimals),
	}
	saveRunesActivityErr := db.RunesActivity.StoreRunesActivity(runesActivity)
	if saveRunesActivityErr != nil {
		log.Error("store RunesActivity error", "error", saveRunesActivityErr)
		return saveRunesActivityErr
	}
	return nil
}

func GetTokenNameAndSymbol(tokenAddress common.Address, ethClient *ethclient.Client) (string, string, uint8, error) {
	token, err := abi.NewIERC20(tokenAddress, ethClient)
	if err != nil {
		return "", "", 0, err
	}
	name, err := token.Name(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, err
	}
	symbol, err := token.Symbol(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, err
	}
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		return "", "", 0, err
	}
	return name, symbol, decimals, nil
}
