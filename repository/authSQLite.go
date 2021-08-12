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
	query := fmt.Sprintf("INSERT INTO %s (fio, login, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.sql.QueryRow(query, user.Fio, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthSQLite) GetUser(login, password string) (mybank.User, error) {
	var user mybank.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE login=$1 AND password_hash=$2", usersTable)
	err := r.db.sql.QueryRow(query, user.Id).Scan(&user.Login, &user.Password)

	return user, err
}
