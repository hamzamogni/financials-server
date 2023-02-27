package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountController struct {
	Controller

	CurrencyRepository *repository.CurrencyRepository
	AccountRepository  *repository.AccountRepository
}

func NewAccountController(currencyRepository *repository.CurrencyRepository, accountRepository *repository.AccountRepository) *AccountController {
	return &AccountController{Controller: Controller{}, CurrencyRepository: currencyRepository, AccountRepository: accountRepository}
}

func (ac AccountController) IndexAccount(c *gin.Context) {
	accounts, err := usecase.NewManageAccount(
		repository.NewAccountRepository(),
		repository.NewCurrencyRepository(),
	).Index()

	if err != nil {
		ac.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ac.response.Success(c, http.StatusOK, accounts)
}

func (ac AccountController) CreateAccount(c *gin.Context) {
	var args usecase.CreateAccountArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		ac.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	account, err := usecase.NewManageAccount(
		repository.NewAccountRepository(),
		repository.NewCurrencyRepository(),
	).Create(args)

	if err != nil {
		ac.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ac.response.Success(c, http.StatusCreated, account)
}
