package controller

import (
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/adapter/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
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
