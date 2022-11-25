package controller

import (
	"errors"
	"financials/internal/app/adapter/repository"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var (
	currencyRepository = repository.Currency{}
)

func (ctrl Controller) IndexCurrency(c *gin.Context) {
	var args usecase.IndexCurrencyArgs

	args.CurrencyRepository = currencyRepository

	currencies, err := usecase.IndexCurrency(args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, currencies)
}

func (ctrl Controller) GetCurrency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	currency, err := usecase.GetCurrency(usecase.GetCurrencyArgs{
		ID:                 uint(id),
		CurrencyRepository: currencyRepository,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, currency)
}

func (ctrl Controller) CreateCurrency(c *gin.Context) {
	var args usecase.CreateCurrencyArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	args.CurrencyRepository = currencyRepository

	currency, err := usecase.CreateCurrency(args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, currency)
}

func (ctrl Controller) UpdateCurrency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad id"})
	}

	var args usecase.UpdateCurrencyArgs
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
	}

	args.ID = uint(id)
	args.CurrencyRepository = currencyRepository

	currency, err := usecase.UpdateCurrency(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, currency)

}

func (ctrl Controller) DeleteCurrency(c *gin.Context) {
	id := c.Param("id")

	err := usecase.DeleteCurrency(usecase.DeleteCurrencyArgs{
		ID:                 id,
		CurrencyRepository: currencyRepository,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
