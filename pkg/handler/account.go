package handler

import "github.com/gin-gonic/gin"

//пополнение счета
func (h *Handler) AddFunds(c *gin.Context) {

}

//расчет процента и внесение его на счет
func (h *Handler) SumProfit(c *gin.Context) {

}

//списания со счета
func (h *Handler) Withdraw(c *gin.Context) {

}

//вывод валюты счета
func (h *Handler) GetCurrency(c *gin.Context) {

}

//вывод курса валюты счета к передаваемой валюте
func (h *Handler) GetAccountCurrencyRate(c *gin.Context) {

}

//вывод баланс счета в указанной валюте
func (h *Handler) GetBalance(c *gin.Context) {

}