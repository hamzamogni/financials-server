package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/adapter/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	response response.Response
}

var (
	currencyRepository = repository.Currency{}
	accountRepository  = repository.Account{}
)

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := Controller{}

	// currencies routes
	r.GET("/currencies", ctrl.IndexCurrency)
	r.GET("/currencies/:id", ctrl.GetCurrency)
	r.POST("/currencies", ctrl.CreateCurrency)
	r.PATCH("/currencies/:id", ctrl.UpdateCurrency)
	r.DELETE("/currencies/:id", ctrl.DeleteCurrency)

	r.POST("/accounts", ctrl.CreateAccount)
	return r
}
