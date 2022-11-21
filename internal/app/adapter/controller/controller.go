package controller

import "github.com/gin-gonic/gin"

type Controller struct{}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	r.POST("/currencies", ctrl.CreateCurrency)
	r.GET("/currencies/:id", ctrl.GetCurrency)

	return r
}
