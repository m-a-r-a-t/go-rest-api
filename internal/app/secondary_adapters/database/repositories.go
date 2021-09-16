package database

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/m-a-r-a-t/go-rest-api/internal/app/core"
	"log"
	"strings"
)

type UserRepo struct {
	Db *sqlx.DB
}

//func (p UserRepo) GetAllUsers() []gin.H {
//
//	rows, err := p.Db.Query(`SELECT * FROM "User"`)
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//	var users []User
//
//	for rows.Next() {
//		u := User{}
//		err := rows.Scan(&u.id, &u.name, &u.age)
//		if err != nil {
//			fmt.Println(err)
//			continue
//		}
//
//		users = append(users, u)
//	}
//
//	var res []gin.H
//
//	//fmt.Println(users[0])
//	for _, value := range users {
//		res = append(res, gin.H{
//			"id":   value.id,
//			"name": value.name,
//			"age":  value.age,
//		})
//	}
//	return res
//}

func (p UserRepo) GetUserBalanceById(id string) (float64, error) {
	var balance float64
	err := p.Db.Get(&balance,
		`SELECT balance FROM "User" WHERE id=$1`,
		id,
	)

	if err != nil {
		str_err := strings.ReplaceAll(err.Error(), " ", "_")
		if str_err == "sql:_no_rows_in_result_set" {
			return balance, errors.New("app.user_not_found")
		}
		return balance, err
	}

	return balance, nil

}

func (p UserRepo) SetUserBalance(id string, value float64) sql.Result {

	result := p.Db.MustExec(
		`INSERT INTO "User" (id,balance) VALUES ($1,$2)
    ON CONFLICT(id) DO UPDATE SET  balance =  $2`,
		id,
		value,
	)

	return result
}

func (p UserRepo) GetUserById(id string) (core.User, error) {

	var user core.User
	err := p.Db.Get(&user,
		`SELECT * FROM "User" WHERE id=$1`,
		id,
	)

	if err != nil {
		str_err := strings.ReplaceAll(err.Error(), " ", "_")
		if str_err == "sql:_no_rows_in_result_set" {
			return user, errors.New("app.user_not_found")
		}
		return user, err
	}

	return user, nil

}

func (p UserRepo) SetUsersBalances(users []core.User) sql.Result {

	result, err := p.Db.NamedExec(
		`INSERT INTO "User" (id,balance) VALUES (:id,:balance)
    ON CONFLICT(id) DO UPDATE SET  balance = excluded.balance`,
		users,
	)
	log.Println(err)

	return result
}
