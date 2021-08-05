package mybank

type users struct {
	id int `json:"id"`
	fio  string `json:"fio"`
	login string `json:"login"`
	password string `json:"password"`
}
