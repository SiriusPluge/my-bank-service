package service

import (
	mybank "github.com/SiriusPluge/my-bank-service"
	"github.com/SiriusPluge/my-bank-service/repository"
)

type Authorization interface {
	CreateUser(user mybank.User) (int, error)
	GenerateToken(login, password string) (string, error)
}

type AccountInterface interface {
	// AddFunds Позволяет внести на счёт сумму sum
	AddFunds(sum float64) error
	// SumProfit Рассчитывает процент по вкладу и полученные деньги вносит на счёт
	SumProfit() error
	// Withdraw Производит списание со счёта по указанным правилам. Если списание выходит за рамки правил, выдаёт ошибку
	Withdraw(sum float64) error
	// GetCurrency Выдаёт валюту счёта
	GetCurrency() (string, error)
	// GetAccountCurrencyRate Выдаёт курс валюты счёта к передаваемой валюте cur
	GetAccountCurrencyRate(cur string) (float64, error)
	// GetBalance Выдаёт баланс счёта в указанной валюте
	GetBalance(cur string) (float64, error)
}

type Service struct {
	Authorization
	AccountInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
