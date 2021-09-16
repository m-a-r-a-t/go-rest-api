package gin_answer

import "github.com/gin-gonic/gin"

func SendResult(c *gin.Context, result gin.H) {
	c.JSON(200, result)
}

func SendError(c *gin.Context, err gin.H) {
	c.JSON(400, err)
}
