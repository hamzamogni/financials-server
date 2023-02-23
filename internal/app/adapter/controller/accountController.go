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
		ctrl.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	account, err := usecase.ManageAccount{
		AccountRepository:  accountRepository,
		CurrencyRepository: currencyRepository,
	}.Create(args)

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusCreated, account)
}
