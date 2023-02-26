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

func (ctrl Controller) IndexCurrency(c *gin.Context) {

	currencies, err := usecase.NewManageCurrency(repository.NewCurrencyRepository()).Index()

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusOK, currencies)
}

func (ctrl Controller) GetCurrency(c *gin.Context) {
	symbol := strings.ToUpper(c.Param("symbol"))

	args := usecase.GetCurrencyArgs{
		Symbol: symbol,
	}

	currency, err := usecase.NewManageCurrency(repository.NewCurrencyRepository()).Get(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctrl.response.Error(c, http.StatusNotFound, err)
			return
		}

		ctrl.response.Error(c, http.StatusInternalServerError, err)
		return
	}

	ctrl.response.Success(c, http.StatusOK, currency)
}

func (ctrl Controller) CreateCurrency(c *gin.Context) {
	var args usecase.CreateCurrencyArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		ctrl.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	currency, err := usecase.NewManageCurrency(
		repository.NewCurrencyRepository(),
	).Create(args)

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusCreated, currency)
}

func (ctrl Controller) DeleteCurrency(c *gin.Context) {
	symbol := cases.Upper(language.Und).String(c.Param("symbol"))

	args := usecase.DeleteCurrencyArgs{
		Symbol: symbol,
	}

	err := usecase.NewManageCurrency(
		repository.NewCurrencyRepository(),
	).Delete(args)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctrl.response.Error(c, http.StatusNotFound, err)
			return
		}

		ctrl.response.Error(c, http.StatusInternalServerError, err)
		return
	}

	ctrl.response.Success(c, http.StatusOK, "")
}
