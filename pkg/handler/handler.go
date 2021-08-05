package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp) // регистрация
		auth.POST("/sign-in", h.signIn) // авторизация
	}

	api := router.Group("/api")
	{
		checks := api.Group("/checks")
		{
			checks.PUT("/", h.AddFunds)               //пополнение счета
			checks.PUT("/", h.SumProfit)              //расчет процента и внесение его на счет
			checks.PUT("/", h.Withdraw)               //списания со счета
			checks.GET("/", h.GetCurrency)            //вывод валюты счета
			checks.GET("/", h.GetAccountCurrencyRate) //вывод курса валюты счета к передаваемой валюте
			checks.GET("/", h.GetBalance)             //вывод баланс счета в указанной валюте
		}
	}

	return router
}
