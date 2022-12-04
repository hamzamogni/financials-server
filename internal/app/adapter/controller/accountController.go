package controller

import (
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl Controller) CreateAccount(c *gin.Context) {
	var args usecase.CreateAccountArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	account, err := usecase.ManageAccount{
		AccountRepository:  accountRepository,
		CurrencyRepository: currencyRepository,
	}.Create(args)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, account)
}
