package repository

import (
	"database/sql"
	_ "database/sql"
	"log"
)

const (
	accountTable      = "account"
	usersTable        = "users"
	currencyTable     = "currency"
	transactionsTable = "transactions"

	insertSQL = `
		INSERT INTO account (
			name, userId
		) VALUES (
			?, ?
		);
		
		INSERT INTO users (
			fio, login, password
		) VALUES (
			?, ?, ?
		);
		
		INSERT INTO currency (
			name, price
		) VALUES (
			?, ?
		);

		INSERT INTO transactions (
			typeCurrency, count, currencyId, typeTransaction
		) VALUES (
			?, ?, ?, ?
		)
	`
	schemaSQL = `
		CREATE TABLE IF NOT EXISTS account (
			id    INT PRIMARY KEY,
			name TEXT,
			userId INT,
		);

		CREATE TABLE IF NOT EXISTS users (
			id    INT PRIMARY KEY,
			fio  TEXT,
			login TEXT,
			password TEXT,
		);

		CREATE TABLE IF NOT EXISTS currency (
			id    INTEGER PRIMARY KEY,
			name TEXT,
			price FLOAT,
		);

		CREATE TABLE IF NOT EXISTS transactions (
			id    INTEGER PRIMARY KEY,
			date DATETIME,
			typeCurrency TEXT,
			count FLOAT,
			currencyId INT,
			typeTransaction INT,
		)
	`
)

type DB struct {
	sql  *sql.DB
	stmt *sql.Stmt
}

func NewDB() (*DB, error) {

	sqlDB, err := sql.Open("sqlite3", "./db/myService.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = sqlDB.Exec(schemaSQL); err != nil {
		return nil, err
	}

	stmt, err := sqlDB.Prepare(insertSQL)
	if err != nil {
		return nil, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	db := DB{
		sql:  sqlDB,
		stmt: stmt,
	}

	return &db, nil
}
