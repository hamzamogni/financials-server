package controller

import "github.com/gin-gonic/gin"

type Controller struct{}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	// currencies routes
	r.GET("/currencies", ctrl.IndexCurrency)
	r.GET("/currencies/:id", ctrl.GetCurrency)
	r.POST("/currencies", ctrl.CreateCurrency)
	r.PATCH("/currencies/:id", ctrl.UpdateCurrency)
	r.DELETE("/currencies/:id", ctrl.DeleteCurrency)

	return r
}
