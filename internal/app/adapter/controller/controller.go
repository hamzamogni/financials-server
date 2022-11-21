package controller

import "github.com/gin-gonic/gin"

type Controller struct{}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	r.GET("/currencies", ctrl.IndexCurrency)
	r.GET("/currencies/:id", ctrl.GetCurrency)
	r.POST("/currencies", ctrl.CreateCurrency)

	return r
}
