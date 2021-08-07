package repository

import (
	"database/sql"
	_ "database/sql"
	"log"
)

func NewDB() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./db/myService.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}


//schemaSQL := `
//	CREATE TABLE IF NOT EXISTS account (
//	id    INT PRIMARY KEY AUTOINCREMENT UNIQUE,
//	name TEXT,
//	userId INT
//	);
//
//	CREATE TABLE IF NOT EXISTS users(
//	id    INT PRIMARY KEY AUTOINCREMENT UNIQUE,
//	fio  TEXT,
//	login TEXT,
//	password TEXT
//	);
//
//	CREATE TABLE IF NOT EXISTS currency(
//	id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
//	name TEXT,
//	price FLOAT,
//	);
//
//	CREATE TABLE IF NOT EXISTS transactions(
//	id    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
//	date DATETIME,
//	typeCurrency TEXT,
//	count FLOAT,
//	currencyId INT,
//	typeTransaction INT
//	)
//`
//
//tx, err := db.Begin()
//if err != nil {
//	log.Fatal(err)
//}


