package app

import (
	"crypto/sha256"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"strconv"
	"time"
)

type Auth struct {
	ID   uint
	User *User

	AccessToken  string
	RefreshToken string
	Expiry       time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type claims struct {
	Username string `json:"username"`
	jwt.MapClaims
}

// JWTOutput is the output of the GenerateToken method
type JWTOutput struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (a *Auth) GenerateToken() error {
	expirationTime := time.Now().Add(100 * time.Minute)

	//claimsValue := &claims{
	//	Username: a.User.Username,
	//	MapClaims: jwt.MapClaims{
	//		"exp": expirationTime.Unix(),
	//	},
	//}

	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsValue)
	//
	//fmt.Println(os.Getenv("JWT_SECRET"))

	//tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	//if err != nil {
	//	return err
	//}

	a.AccessToken = strconv.Itoa(int(a.User.ID)) + "|" + randomString(40)
	a.Expiry = expirationTime

	return nil
}

func randomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUBWXYZ"
	x := make([]byte, length)
	for i := range x {
		x[i] = letters[rand.Int63()%int64(len(letters))]
	}

	ret := fmt.Sprintf("%+x", hash(x))

	fmt.Println(x)
	fmt.Println(ret)
	return ret
}

func hash(str []byte) []byte {
	h := sha256.New()
	h.Write(str)
	bs := h.Sum(nil)

	return bs
}

type AuthService interface {
	FindAuthByID(id uint) (*Auth, error)

	CreateAuth(auth *Auth) (*Auth, error)

	DeleteAuth(id uint) error
}
