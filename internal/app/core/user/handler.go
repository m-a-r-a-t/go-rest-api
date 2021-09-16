package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/core"
	"github.com/m-a-r-a-t/go-rest-api/pkg/pg_database"
	"log"
)

type GetUserBalanceHandler struct {
	UserRepo              IUserRepo
	ExchangesRatesApiRepo IExchangesRatesApiRepo
}

func (h GetUserBalanceHandler) Handle(c *gin.Context) (gin.H, gin.H) {
	var params struct {
		id       string
		currency string
	}
	params.id, params.currency = c.Query("id"), c.Query("currency")

	balance, err := h.UserRepo.GetUserBalanceById(params.id)
	if err != nil {
		e := UserNotFound{
			id: params.id,
		}
		return nil, e.Error()
	}

	if params.currency != "" {
		convertedBalance, err := h.ExchangesRatesApiRepo.Convert(
			"RUB",
			params.currency,
			balance,
		)
		if err != nil {
			log.Println(err)
			e := CurrencyConverterApiError{}
			return nil, e.Error()
		}
		return gin.H{"balance": convertedBalance}, nil

	}

	return gin.H{"balance": balance}, nil
}

type AddOrWithdrawFunds struct {
	UserRepo IUserRepo
}

func (h AddOrWithdrawFunds) Handle(c *gin.Context) (gin.H, gin.H) {
	var params struct {
		UserId     string
		Funds      float64
		IsWithdraw bool
	}
	bytes, _ := c.GetRawData()
	_ = json.Unmarshal(bytes, &params)
	var newBalance float64

	if params.UserId == "" {
		e := UserNotFound{}
		return nil, e.Error()
	} else if params.Funds <= 0 {
		log.Println(params.Funds)
		e := FundsError{}
		return nil, e.Error()
	}

	log.Println("params", params)

	balance, err := h.UserRepo.GetUserBalanceById(params.UserId)
	if params.IsWithdraw && err != nil {
        log.Println(err)
		e := UserNotFound{id: params.UserId}
		return nil, e.Error()
	}

	if params.IsWithdraw {
		newBalance = balance - params.Funds
		if newBalance < 0 {
			e := WithdrawError{}
			return nil, e.Error()
		}

	} else {
		newBalance = balance + params.Funds
	}

	log.Println(newBalance)

	sqlResult := h.UserRepo.SetUserBalance(params.UserId, newBalance)
	log.Println(sqlResult)
	if sqlResult == nil {
		e := pg_database.DatabaseError{}
		return nil, e.Error()
	}
	return gin.H{"success": true}, nil
}

type TransferFunds struct {
	UserRepo IUserRepo
}

func (h TransferFunds) Handle(c *gin.Context) (gin.H, gin.H) {
	var params struct {
		FromUserId string
		ToUserId   string
		Funds      float64
	}

	bytes, _ := c.GetRawData()
	_ = json.Unmarshal(bytes, &params)
	var newBalances []core.User

	if params.FromUserId == "" || params.ToUserId == "" {
		e := UserNotFound{}
		return nil, e.Error()
	} else if params.Funds <= 0 {
		log.Println(params.Funds)
		e := FundsError{}
		return nil, e.Error()
	}

	log.Println("params", params)

	senderMoneyBalance, err := h.UserRepo.GetUserBalanceById(params.FromUserId)
	if err != nil {
		e := UserNotFound{id: params.FromUserId}
		return nil, e.Error()
	}

	log.Println("Sender money ", senderMoneyBalance)

	getterMoneyBalance, err := h.UserRepo.GetUserBalanceById(params.ToUserId)
	if err != nil {
		getterMoneyBalance = 0.0
	}

	log.Println("Getter money ", getterMoneyBalance)

	newSenderBalance := core.User{
		Id:      params.FromUserId,
		Balance: senderMoneyBalance - params.Funds,
	}

	if newSenderBalance.Balance < 0 {
		e := WithdrawError{}
		return nil, e.Error()
	}

	newGetterBalance := core.User{
		Id:      params.ToUserId,
		Balance: getterMoneyBalance + params.Funds,
	}

	newBalances = append(newBalances, newSenderBalance, newGetterBalance)

	sqlResult := h.UserRepo.SetUsersBalances(newBalances)
	log.Println(sqlResult)
	if sqlResult == nil {
		e := pg_database.DatabaseError{}
		return nil, e.Error()
	}
	return gin.H{"success": true}, nil
}

//insertDynStmt := `insert into "User"("id", "name","age") values($1, $2,$3)`
//_, e := b.Db.Exec(insertDynStmt, 1, "Marat", 20)
//pg_database.CheckError(e)
