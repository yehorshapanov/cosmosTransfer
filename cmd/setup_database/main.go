package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "mysecretpassword"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	createBlocksTable(db)
	createTransactionsTable(db)
	createTransfersTable(db)
	createView(db)
}

func createTransfersTable(db *sql.DB) {
	sqlStatement := `
CREATE TABLE transfers (
 id SERIAL PRIMARY KEY,
 transaction_hash TEXT,
 from_where TEXT,
 to_where TEXT,
 amount INT,
 created_at TIMESTAMPTZ
)`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}

func createTransactionsTable(db *sql.DB) {
	sqlStatement := `
CREATE TABLE transactions (
 id SERIAL PRIMARY KEY,
 block_number INT,
 transaction_hash TEXT,
 fee INT,
 payer_address TEXT,
 created_at TIMESTAMPTZ
)`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}

func createBlocksTable(db *sql.DB) {
	sqlStatement := `
CREATE TABLE blocks (
 id SERIAL PRIMARY KEY,
 block_number INT,
 block_hash TEXT,
 block_proposer_address TEXT,
 created_at TIMESTAMPTZ
)`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}

}

func createView(db *sql.DB) {
	sqlStatement := `
CREATE VIEW vwTransferTransaction AS
 SELECT b.block_number,
    t.transaction_hash,
    tf.amount,
    tf.from_where,
    tf.to_where
   FROM blocks b
     JOIN transactions t ON b.block_number = t.block_number
     JOIN transfers tf ON t.transaction_hash = tf.transaction_hash;`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}
