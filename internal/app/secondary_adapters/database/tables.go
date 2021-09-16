package database

import "github.com/jmoiron/sqlx"

type User struct {
    id   int
    balance  int64
}



var userShema = `
CREATE TABLE IF NOT EXISTS "User" (
    id text unique not null primary key,
    balance float not null 
);`




func CreateUserTable(db *sqlx.DB){
    db.MustExec(userShema)
}