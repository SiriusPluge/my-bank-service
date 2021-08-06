package main

import (
	"database/sql"
)

const (
	insertSQL = `
INSERT INTO trades (
	time, symbol, price, buy
) VALUES (
	?, ?, ?, ?
)
`

	schemaSQL = `
CREATE TABLE IF NOT EXISTS account (
    id    INT PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT,
    userId INT
);

CREATE TABLE IF NOT EXISTS users(
    id    INT PRIMARY KEY AUTOINCREMENT UNIQUE,
    fio  TEXT,
    login TEXT,
    password TEXT
);

CREATE TABLE IF NOT EXISTS currency(
    id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT,
    price FLOAT,
);

CREATE TABLE IF NOT EXISTS transactions(
    id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
    date DATETIME,
    typeCurrency TEXT,
    count FLOAT,
    currencyId INT,
    typeTransaction INT
)
`
)

type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
}

// NewDB
func NewDB(dbFile string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if _, err = sqlDB.Exec(schemaSQL); err != nil {
		return nil, err
	}

	stmt, err := sqlDB.Prepare(insertSQL)
	if err != nil {
		return nil, err
	}

	db := DB{
		sql:    sqlDB,
		stmt:   stmt,
	}
	return &db, nil
}

//// Add stores a trade into the buffer. Once the buffer is full, the
//// trades are flushed to the database.
//func (db *DB) Add(trade Trade) error {
//	if len(db.buffer) == cap(db.buffer) {
//		return errors.New("trades buffer is full")
//	}
//
//	db.buffer = append(db.buffer, trade)
//	if len(db.buffer) == cap(db.buffer) {
//		if err := db.Flush(); err != nil {
//			return fmt.Errorf("unable to flush trades: %w", err)
//		}
//	}
//
//	return nil
//}
//
//// Flush inserts pending trades into the database.
//func (db *DB) Flush() error {
//	tx, err := db.sql.Begin()
//	if err != nil {
//		return err
//	}
//
//	for _, trade := range db.buffer {
//		_, err := tx.Stmt(db.stmt).Exec(trade.Time, trade.Symbol, trade.Price, trade.IsBuy)
//		if err != nil {
//			tx.Rollback()
//			return err
//		}
//	}
//
//	db.buffer = db.buffer[:0]
//	return tx.Commit()
//}
//
//// Close flushes all trades to the database and prevents any future trading.
//func (db *DB) Close() error {
//	defer func() {
//		db.stmt.Close()
//		db.sql.Close()
//	}()
//
//	if err := db.Flush(); err != nil {
//		return err
//	}
//
//	return nil
//}




