package http

import (
	"financials/internal/app"
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

type SignInArgs struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) SignIn(c *gin.Context) {
	var args SignInArgs

	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.UserService.GetByEmail(args.Email)
	if err != nil {
		s.ErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	if user.Password != args.Password {
		s.ErrorResponse(c, http.StatusUnauthorized, err)
		return
	}

	auth := app.Auth{
		User: user,
	}
	if err := auth.GenerateToken(); err != nil {
		s.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	newAuth, err := s.AuthService.CreateAuth(&auth)
	if err != nil {
		s.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	s.SuccessResponse(c, http.StatusOK, newAuth)
}
