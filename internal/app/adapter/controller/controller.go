package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/adapter/response"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	response response.Response
}

func Router() *gin.Engine {
	r := gin.Default()

	accountController := NewAccountController(repository.NewCurrencyRepository(), repository.NewAccountRepository())
	currencyController := NewCurrencyController(repository.NewCurrencyRepository())
	authController := NewAuthController()

	r.POST("/signin", authController.SignIn)

	// currencies routes
	r.Use(authController.AuthMiddleware()).GET("/currencies", currencyController.IndexCurrency)
	r.GET("/currencies/:symbol", currencyController.GetCurrency)
	r.POST("/currencies", currencyController.CreateCurrency)
	r.DELETE("/currencies/:symbol", currencyController.DeleteCurrency)

	r.GET("/accounts", accountController.IndexAccount)
	r.POST("/accounts", accountController.CreateAccount)
	return r
}
