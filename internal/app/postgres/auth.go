package postgres

import (
	"financials/internal/app"
	"gorm.io/gorm"
	"time"
)

type Auth struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User

	AccessToken  string
	RefreshToken string
	Expiry       time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthService struct {
	db *gorm.DB
}

func NewAuthService() *AuthService {
	return &AuthService{db: Connection()}
}

func (a AuthService) FindAuthByID(id uint) (*app.Auth, error) {
	var auth Auth

	result := a.db.Joins("User", a.db.Where(&Auth{ID: id})).First(&auth)
	if result.Error != nil {
		return &app.Auth{}, result.Error
	}

	return authFromModel(&auth), nil
}

func (a AuthService) CreateAuth(auth *app.Auth) (*app.Auth, error) {
	newAuth, err := createAuth(a.db, auth)
	if err != nil {
		return nil, err
	}

	return newAuth, nil
}

func (a AuthService) DeleteAuth(id uint) error {
	//TODO implement me
	panic("implement me")
}

func createAuth(db *gorm.DB, auth *app.Auth) (*app.Auth, error) {
	now := time.Now()
	newAuth := Auth{
		User:         User{ID: auth.User.ID},
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		Expiry:       auth.Expiry,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if result := db.Create(&newAuth); result.Error != nil {
		return nil, result.Error
	}

	return authFromModel(&newAuth), nil
}

func authFromModel(auth *Auth) *app.Auth {
	return &app.Auth{
		ID: auth.ID,
		User: &app.User{
			Username: auth.User.Username,
			Email:    auth.User.Email,
			Password: auth.User.Password,
		},
		AccessToken:  auth.AccessToken,
		RefreshToken: auth.RefreshToken,
		Expiry:       auth.Expiry,
		CreatedAt:    auth.CreatedAt,
		UpdatedAt:    auth.UpdatedAt,
	}
}
