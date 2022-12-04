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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(200, currencies)
}

func (ctrl Controller) GetCurrency(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	args := usecase.GetCurrencyArgs{
		Id: uint(id),
	}

	currency, err := usecase.ManageCurrency{CurrencyRepository: currencyRepository}.Get(args)
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

	currency, err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Create(args)

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

	args.Id = uint(id)

	currency, err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Update(args)

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

	args := usecase.DeleteCurrencyArgs{
		Id: id,
	}

	err := usecase.ManageCurrency{
		CurrencyRepository: currencyRepository,
	}.Delete(args)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
