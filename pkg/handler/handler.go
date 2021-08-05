package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/users")
	{
		auth.POST("/registration", h.registration) // регистрация
		auth.POST("/auth", h.signIn)               // авторизация
	}

	api := router.Group("/api")
	{
		checks := api.Group("/checks")
		{
			checks.PUT("/", h.AddFunds)                                     //пополнение счета
			checks.PUT("/sumProfit", h.SumProfit)                           //расчет процента и внесение его на счет
			checks.PUT("/withdraw", h.Withdraw)                             //списания со счета
			checks.GET("/getCurrency", h.GetCurrency)                       //вывод валюты счета
			checks.GET("/getAccountCurrencyRate", h.GetAccountCurrencyRate) //вывод курса валюты счета к передаваемой валюте
			checks.GET("/getBalance", h.GetBalance)                         //вывод баланс счета в указанной валюте
		}
	}

	return router
}
