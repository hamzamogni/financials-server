package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct{}

func (r Response) Success(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func (r Response) Error(c *gin.Context, statusCode int, message error) {
	c.JSON(statusCode, gin.H{
		"error": message.Error(),
	})
}
