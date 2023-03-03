package http

import (
	"financials/internal/app/http/service"
	"financials/internal/app/postgres"
	"financials/internal/app/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) registerAuthRoutes(r *gin.Engine) {
	r.POST("/signin", s.SignIn)
	r.POST("/signup", s.SignUp)
}

func (s *Server) SignUp(c *gin.Context) {
	var args usecase.CreateUserArgs
	err := c.ShouldBindJSON(&args)
	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := usecase.NewManageUser(postgres.NewUserService()).Create(args)
	if err != nil {
		s.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	s.SuccessResponse(c, http.StatusCreated, user)
}

func (s *Server) SignIn(c *gin.Context) {
	var user postgres.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Username != "hmogni" || user.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := service.AuthService{}.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, token)
}
