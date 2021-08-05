package mybank

type check struct {
	id int `json:"id"`
	name string `json:"name"`
	userId int `json:"user_id"`
}

type currency struct {
	id int `json:"id"`
	name string `json:"name"`
	price float32 `json:"price"`
}

type transactions struct {
	id int `json:"id"`
	typeCurrency string `json:"type_currency"`
	count float32 `json:"count"`
	currencyId int `json:"count"`
	typeTransaction int `json:"type_transaction"`
}
