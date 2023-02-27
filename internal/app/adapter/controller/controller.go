package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/adapter/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	response response.Response

	CurrencyRepository repository.CurrencyRepository
	AccountRepository  repository.AccountRepository
}

func NewController(
	currencyRepository repository.CurrencyRepository,
	accountRepository repository.AccountRepository,
) *Controller {
	return &Controller{
		response:           response.Response{},
		CurrencyRepository: currencyRepository,
		AccountRepository:  accountRepository,
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	ctrl := NewController(*repository.NewCurrencyRepository(), *repository.NewAccountRepository())

	r.POST("/signin", ctrl.SignIn)
	// currencies routes
	r.Use(ctrl.AuthMiddleware()).GET("/currencies", ctrl.IndexCurrency)
	r.GET("/currencies/:symbol", ctrl.GetCurrency)
	r.POST("/currencies", ctrl.CreateCurrency)
	r.DELETE("/currencies/:symbol", ctrl.DeleteCurrency)

	r.GET("/accounts", ctrl.IndexAccount)
	r.POST("/accounts", ctrl.CreateAccount)
	return r
}
