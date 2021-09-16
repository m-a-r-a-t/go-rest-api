package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/m-a-r-a-t/go-rest-api/pkg/gin_answer"
	"log"
)

type HandleFunc func(c *gin.Context) (gin.H, gin.H)

//type

type Handler interface {
	Handle(c *gin.Context) (gin.H, gin.H)
}

func CustomDecorator(fn HandleFunc) func(c *gin.Context) {

	return func(c *gin.Context) {
		log.Println("Начало исполнения с целым числом")
		result, err := fn(c)

		if err != nil {
			gin_answer.SendError(c, err)
			return
		}

		gin_answer.SendResult(c, result)
		log.Println("Исполнение завершено с результатом")

	}
}

type Error struct {
	Msg         string
	TypeOfError string
}

type CustomError interface {
    Error() string
}