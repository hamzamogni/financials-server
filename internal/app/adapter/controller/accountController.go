package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctrl Controller) IndexAccount(c *gin.Context) {
	accounts, err := usecase.NewManageAccount(
		repository.NewAccountRepository(),
		repository.NewCurrencyRepository(),
	).Index()

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusOK, accounts)
}

func (ctrl Controller) CreateAccount(c *gin.Context) {
	var args usecase.CreateAccountArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		ctrl.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	account, err := usecase.NewManageAccount(
		repository.NewAccountRepository(),
		repository.NewCurrencyRepository(),
	).Create(args)

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusCreated, account)
}
