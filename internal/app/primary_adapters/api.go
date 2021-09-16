package primary_adapters

import (
	"github.com/gin-gonic/gin"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/core/user"
	"github.com/m-a-r-a-t/go-rest-api/pkg/utils"
)

func Routes(
	getUserBalanceHandler user.GetUserBalanceHandler,
	addOrWithdrawFunds user.AddOrWithdrawFunds,
	transferFundsHandler user.TransferFunds,
) {
	r := gin.Default()

	r.GET("/get_user_balance", utils.CustomDecorator(getUserBalanceHandler.Handle))

	r.POST("/add_or_withdraw_funds", utils.CustomDecorator(addOrWithdrawFunds.Handle))
	r.POST("/transfer_funds", utils.CustomDecorator(transferFundsHandler.Handle))

	r.Run()
}
