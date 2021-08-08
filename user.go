package mybank

type User struct {
	Id       int    `json:"-" db:"id"`
	Fio      string `json:"fio" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}
