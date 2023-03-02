package controller

import (
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/adapter/repository"
	"financials/internal/app/adapter/service"
	"financials/internal/app/application/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	Controller
}

func NewAuthController() *AuthController {
	return &AuthController{Controller: Controller{}}
}

func (ac AuthController) SignUp(c *gin.Context) {
	var args usecase.CreateUserArgs
	err := c.ShouldBindJSON(&args)
	if err != nil {
		ac.response.Error(c, http.StatusBadRequest, err)
		return
	}

	user, err := usecase.NewManageUser(repository.NewUserRepository()).Create(args)
	if err != nil {
		ac.response.Error(c, http.StatusBadRequest, err)
		return
	}

	ac.response.Success(c, http.StatusCreated, user)
}

func (ac AuthController) SignIn(c *gin.Context) {
	var user model.User

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

func (ac AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")

		err := service.AuthService{}.ValidateToken(tokenValue)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
