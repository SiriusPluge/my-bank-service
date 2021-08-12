package mybank

import "errors"

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type User struct {
	Id       int    `json:"-" db:"id"`
	Fio      string `json:"fio" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
