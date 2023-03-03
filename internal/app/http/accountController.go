package http

import (
	"financials/internal/app/postgres"
	"financials/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) registerAccountRoutes(r *gin.Engine) {
	r.GET("/accounts", s.IndexAccount)
	r.POST("/accounts", s.CreateAccount)
}

func (s *Server) IndexAccount(c *gin.Context) {
	accounts, err := usecase.NewManageAccount(
		postgres.NewAccountService(),
		postgres.NewCurrencyService(),
	).Index()

	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	s.SuccessResponse(c, http.StatusOK, accounts)
}

func (s *Server) CreateAccount(c *gin.Context) {
	var args usecase.CreateAccountArgs

	err := c.ShouldBindJSON(&args)
	if err != nil {
		s.ErrorResponse(c, http.StatusUnprocessableEntity, err)
		return
	}

	account, err := usecase.NewManageAccount(
		postgres.NewAccountService(),
		postgres.NewCurrencyService(),
	).Create(args)

	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	s.SuccessResponse(c, http.StatusCreated, account)
}
