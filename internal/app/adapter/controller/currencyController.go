package controller

import (
	"errors"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func (ctrl Controller) IndexCurrency(c *gin.Context) {

	currencies, err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Index()

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusOK, currencies)
}

func (ctrl Controller) GetCurrency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	args := usecase.GetCurrencyArgs{
		Id: uint(id),
	}

	currency, err := usecase.ManageCurrency{CurrencyRepository: currencyRepository}.Get(args)
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

	currency, err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Create(args)

	if err != nil {
		ctrl.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ctrl.response.Success(c, http.StatusCreated, currency)
}

func (ctrl Controller) UpdateCurrency(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var args usecase.UpdateCurrencyArgs
	if err := c.ShouldBindJSON(&args); err != nil {
		ctrl.response.Error(c, http.StatusUnprocessableEntity, err)
		return
	}

	args.Id = uint(id)

	currency, err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Update(args)

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

func (ctrl Controller) DeleteCurrency(c *gin.Context) {
	id := c.Param("id")

	args := usecase.DeleteCurrencyArgs{
		Id: id,
	}

	err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Delete(args)

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
