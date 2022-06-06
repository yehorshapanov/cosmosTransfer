package main

import (
	"context"
	"cosmosTransfer/app"
	"cosmosTransfer/app/cosmos_api"
	"cosmosTransfer/app/db_manager"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func connectionInfo() string {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

var manager *db_manager.DBManager

func main() {
	db, err := sql.Open("postgres", connectionInfo())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	manager = db_manager.NewManager(db)

	c := cosmos_api.NewClient()
	const startValue = 9989379
	for i := startValue; i < startValue+10; i++ {
		ProcessBlock(c, i)
	}
}

func ProcessBlock(c *cosmos_api.Client, blockNumber int) {
	fmt.Printf("------------- Start block %d -------------\n", blockNumber)
	ctx := context.Background()
	block, _ := c.GetBlock(ctx, blockNumber)

	txs := block.Block.Data.Txs

	if txs == nil {
		return
	}

	manager.InsertBlock(block)

	for _, tx := range txs {
		transaction, err := c.GetTx(ctx, app.CalculateTransactionID(tx))
		if err != nil {
			continue
		}

		for _, msg := range transaction.TxDetails.Value.Msg {
			fmt.Printf("------------- Start transaction -------------")
			manager.InsertTransaction(transaction, blockNumber)
			fmt.Printf("%s%s", msg.Type, transaction.Hash)
			if msg.Type == "cosmos-sdk/MsgSend" {
				manager.InsertTransfer(transaction)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("------------- Finish block -------------\n")
}
