package controller

import (
	"financials/internal/app/adapter/repository"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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
	id := c.Param("id")

	currency, err := usecase.GetCurrency(usecase.GetCurrencyArgs{
		ID:                 id,
		CurrencyRepository: currencyRepository,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, currency)
}

func (ctrl Controller) CreateCurrency(c *gin.Context) {
	var args usecase.CreateCurrencyArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
	}

	args.CurrencyRepository = currencyRepository

	currency, err := usecase.CreateCurrency(args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(200, currency)
}
