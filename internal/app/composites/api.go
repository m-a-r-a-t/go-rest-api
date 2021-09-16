package composites

import (
	"github.com/m-a-r-a-t/go-rest-api/internal/app/core/user"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/primary_adapters"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/secondary_adapters/database"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/secondary_adapters/exchangesratesapi"
	"github.com/m-a-r-a-t/go-rest-api/pkg/pg_database"
)

func InitHandlers() {
	s := Settings()
	db := pg_database.Database(s.DatabaseConf)

	defer db.Close()


    database.CreateUserTable(db)


    userRepo := database.UserRepo{
		Db: db,
	}

	exchangesRatesApiRepo := exchangesratesapi.NewExchangesRatesApiRepo()

	getUserBalanceHandler := user.GetUserBalanceHandler{
		UserRepo:              userRepo,
		ExchangesRatesApiRepo: exchangesRatesApiRepo,
	}

	addOrWithdrawFunds := user.AddOrWithdrawFunds{
		UserRepo: userRepo,
	}

	transferFundsHandler := user.TransferFunds{
		UserRepo: userRepo,
	}

	primary_adapters.Routes(
		getUserBalanceHandler,
		addOrWithdrawFunds,
		transferFundsHandler,
	)

}
