package pg_database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DatabaseConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func Database(c DatabaseConfig) *sqlx.DB {
	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Pass,
		c.Name,
	)

	var db *sqlx.DB
	for true {
		db, err := sqlx.Connect("postgres", psqlconn)
		if err != nil {
			log.Println(err)
			log.Println("Waiting postgre")
			time.Sleep(1 * time.Second)
			continue
		}

		err = db.Ping()
		CheckError(err)

		fmt.Println("Connected!")

		return db

	}

	return db

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type DatabaseError struct {
}

func (u *DatabaseError) Error() gin.H {
	return gin.H{
		"type": "app.database_error",
		"msg":  "Wrong with db",
	}
}
