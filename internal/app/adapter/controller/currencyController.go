package controller

import (
	"errors"
	"financials/internal/app/adapter/repository"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type CurrencyController struct {
	Controller
	CurrencyRepository *repository.CurrencyRepository
}

func NewCurrencyController(currencyRepository *repository.CurrencyRepository) *CurrencyController {
	return &CurrencyController{Controller: Controller{}, CurrencyRepository: currencyRepository}
}

func (currencyCtrl CurrencyController) IndexCurrency(c *gin.Context) {

	currencies, err := usecase.NewManageCurrency(repository.NewCurrencyRepository()).Index()

	if err != nil {
		currencyCtrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	currencyCtrl.response.Success(c, http.StatusOK, currencies)
}

func (currencyCtrl CurrencyController) GetCurrency(c *gin.Context) {
	symbol := strings.ToUpper(c.Param("symbol"))

	args := usecase.GetCurrencyArgs{
		Symbol: symbol,
	}

	currency, err := usecase.NewManageCurrency(repository.NewCurrencyRepository()).Get(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			currencyCtrl.response.Error(c, http.StatusNotFound, err)
			return
		}

		currencyCtrl.response.Error(c, http.StatusInternalServerError, err)
		return
	}

	currencyCtrl.response.Success(c, http.StatusOK, currency)
}

func (currencyCtrl CurrencyController) CreateCurrency(c *gin.Context) {
	var args usecase.CreateCurrencyArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		currencyCtrl.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	currency, err := usecase.NewManageCurrency(
		repository.NewCurrencyRepository(),
	).Create(args)

	if err != nil {
		currencyCtrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	currencyCtrl.response.Success(c, http.StatusCreated, currency)
}

func (currencyCtrl CurrencyController) DeleteCurrency(c *gin.Context) {
	symbol := cases.Upper(language.Und).String(c.Param("symbol"))

	args := usecase.DeleteCurrencyArgs{
		Symbol: symbol,
	}

	err := usecase.NewManageCurrency(
		repository.NewCurrencyRepository(),
	).Delete(args)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			currencyCtrl.response.Error(c, http.StatusNotFound, err)
			return
		}

		currencyCtrl.response.Error(c, http.StatusInternalServerError, err)
		return
	}

	currencyCtrl.response.Success(c, http.StatusOK, "")
}
