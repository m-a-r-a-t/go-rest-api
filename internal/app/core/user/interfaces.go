package user

import (
	"database/sql"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/core"
)

type IUserRepo interface {
	//GetAllUsers() []gin.H
	GetUserById(id string) (core.User, error)
	GetUserBalanceById(id string) (float64, error)
	SetUserBalance(id string, value float64) sql.Result
	SetUsersBalances(users []core.User) sql.Result
}

type IExchangesRatesApiRepo interface {
	Convert(
		fromCurrency string,
		toCurrency string,
		amount float64,
	) (float64, error)
}
