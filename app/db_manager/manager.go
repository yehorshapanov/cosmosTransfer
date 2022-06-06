package db_manager

import (
	"cosmosTransfer/app/cosmos_api"
	"database/sql"
)

type DBManager struct {
	db *sql.DB
}

func NewManager(db *sql.DB) *DBManager {
	return &DBManager{db: db}
}

func (d *DBManager) InsertBlock(b *cosmos_api.BlockObject) error {
	sqlStatement := `
INSERT INTO blocks (block_number, block_hash, block_proposer_address, created_at)
VALUES ($1, $2, $3, $4)`
	_, err := d.db.Exec(sqlStatement, b.Block.Header.Height, b.BlockId.Hash, b.Block.ProposerAddress, b.Block.Header.Time)

	return err
}

func (d *DBManager) InsertTransaction(t *cosmos_api.Tx, blockNumber int) error {
	sqlStatement := `
INSERT INTO transactions (block_number, transaction_hash, fee, payer_address, created_at)
VALUES ($1, $2, $3, $4, $5)`
	_, err := d.db.Exec(sqlStatement, blockNumber, t.Hash, t.GetFee(), t.GetFromAddress(), t.Timestamp)

	return err
}

func (d *DBManager) InsertTransfer(t *cosmos_api.Tx) error {
	sqlStatement := `
INSERT INTO transfers (transaction_hash, from_where, to_where, amount, created_at)
VALUES ($1, $2, $3, $4, $5)`
	_, err := d.db.Exec(sqlStatement, t.Hash, t.GetFromAddress(), t.GetToAddress(), t.GetAmount(), t.Timestamp)

	return err
}
