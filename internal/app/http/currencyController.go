package http

import (
	"errors"
	"financials/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func (s *Server) registerCurrencyRouter(r *gin.Engine) {
	r.GET("/currencies", s.IndexCurrency)
	r.GET("/currencies/:symbol", s.GetCurrency)
	r.POST("/currencies", s.CreateCurrency)
	r.DELETE("/currencies/:symbol", s.DeleteCurrency)
}

func (s *Server) IndexCurrency(c *gin.Context) {

	currencies, err := usecase.NewManageCurrency(s.CurrencyService).Index()

	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	s.SuccessResponse(c, http.StatusOK, currencies)
}

func (s *Server) GetCurrency(c *gin.Context) {
	symbol := strings.ToUpper(c.Param("symbol"))

	args := usecase.GetCurrencyArgs{
		Symbol: symbol,
	}

	currency, err := usecase.NewManageCurrency(s.CurrencyService).Get(args)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.ErrorResponse(c, http.StatusNotFound, err)
			return
		}

		s.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	s.SuccessResponse(c, http.StatusOK, currency)
}

func (s *Server) CreateCurrency(c *gin.Context) {
	var args usecase.CreateCurrencyArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		s.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	currency, err := usecase.NewManageCurrency(s.CurrencyService).Create(args)

	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	s.SuccessResponse(c, http.StatusCreated, currency)
}

func (s *Server) DeleteCurrency(c *gin.Context) {
	symbol := cases.Upper(language.Und).String(c.Param("symbol"))

	args := usecase.DeleteCurrencyArgs{
		Symbol: symbol,
	}

	err := usecase.NewManageCurrency(s.CurrencyService).Delete(args)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.ErrorResponse(c, http.StatusNotFound, err)
			return
		}

		s.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	s.SuccessResponse(c, http.StatusOK, "")
}
