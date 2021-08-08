package repository

import (
	"fmt"
	mybank "github.com/SiriusPluge/my-bank-service"
)

type AuthSQLite struct {
	db *DB
}

func NewAuthSQLite(db *DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (r *AuthSQLite) CreateUser(user mybank.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (fio, login, password_hash) values ($1, $2, $3) RETURNING id")

	row := r.db.sql.QueryRow(query, user.Fio, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
