package http

import (
	"financials/internal/app"
	"financials/internal/app/http/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	router *gin.Engine

	AccountService  app.AccountService
	CurrencyService app.CurrencyService
	AuthService     app.AuthService
}

func NewServer() *Server {
	s := &Server{
		router: gin.Default(),
	}

	s.registerAccountRoutes(s.router)
	s.registerCurrencyRouter(s.router)
	s.registerAuthRoutes(s.router)
	return s
}

func (s *Server) SuccessResponse(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func (s *Server) ErrorResponse(c *gin.Context, statusCode int, message error) {
	c.JSON(statusCode, gin.H{
		"error": message.Error(),
	})
}

func (s *Server) Serve() {
	s.router.Run(":8080")
}

// requireAuth is a middleware to validate API key
func (s *Server) requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.GetHeader("Authorization")

		err := service.AuthService{}.ValidateToken(tokenValue)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Next()
	}
}
