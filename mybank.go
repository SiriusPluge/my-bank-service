package mybank

type Account struct {
	id     int    `json:"id"`
	name   string `json:"name"`
	userId int    `json:"user_id"`
}

type Currency struct {
	id    int     `json:"id"`
	name  string  `json:"name"`
	price float32 `json:"price"`
}

type Transactions struct {
	id              int     `json:"id"`
	typeCurrency    string  `json:"type_currency"`
	count           float32 `json:"count"`
	currencyId      int     `json:"count"`
	typeTransaction int     `json:"type_transaction"`
}
