package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// AuthService is a service that handles authentication
type AuthService struct{}

// claims is a custom struct that implements the jwt.Claims interface
type claims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

// JWTOutput is the output of the GenerateToken method
type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

// GenerateToken generates a JWT token
func (as AuthService) GenerateToken(username string) (*JWTOutput, error) {
	expirationTime := time.Now().Add(100 * time.Minute)
	claimsValue := &claims{
		Username: username,
		MapClaims: jwt.MapClaims{
			"exp": expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsValue)

	fmt.Println(os.Getenv("JWT_SECRET"))

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return &JWTOutput{}, err
	}

	return &JWTOutput{
		Token:   tokenString,
		Expires: expirationTime,
	}, nil
}

// ValidateToken validates a JWT token
func (as AuthService) ValidateToken(tokenValue string) error {
	claimsValue := &claims{}
	token, err := jwt.ParseWithClaims(tokenValue, claimsValue, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || token == nil || !token.Valid {
		return err
	}

	return nil
}
