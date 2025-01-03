package create_table

import (
	"fmt"

	"github.com/Pump-Elf-Ranch/per_apollo/database"
)

func CreateTableFromTemplate(chainId string, db *database.DB) {
	createBlockHeaders(chainId, db)
	createContractEvents(chainId, db)
}

func createBlockHeaders(chainId string, db *database.DB) {
	tableName := "template_block_headers"
	tableNameByChainId := fmt.Sprintf("block_headers_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}

func createContractEvents(chainId string, db *database.DB) {
	tableName := "template_contract_events"
	tableNameByChainId := fmt.Sprintf("contract_events_%s", chainId)
	db.CreateTable.CreateTable(tableNameByChainId, tableName)
}
