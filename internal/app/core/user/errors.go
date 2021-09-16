package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserNotFound struct {
	id string
}

func (u *UserNotFound) Error() gin.H {
	return gin.H{
		"type": "app.user_not_found",
		"msg":  fmt.Sprintf("The user with id '%s' was not found.", u.id),
	}
}

type FundsError struct{}

func (u *FundsError) Error() gin.H {
	return gin.H{
		"type": "app.funds_error",
		"msg":  "The amount of funds is lower or equal to zero.",
	}
}

type WithdrawError struct{}

func (u *WithdrawError) Error() gin.H {
	return gin.H{
		"type": "app.withdraw_error",
		"msg":  "Insufficient funds.",
	}
}

type CurrencyConverterApiError struct{}

func (u *CurrencyConverterApiError) Error() gin.H {
	return gin.H{
		"type": "app.currency_converter_api_error",
		"msg":  "Error with currency converter api.",
	}
}
