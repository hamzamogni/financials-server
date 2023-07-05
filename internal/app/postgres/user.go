package postgres

import (
	"financials/internal/app"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint `gorm:"primarykey"`
	Email    string
	Username string
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserService struct {
	Db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{Db: Connection()}
}

func (u UserService) Save(user *app.User) (*app.User, error) {
	newUser, err := createUser(u.Db, user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u UserService) GetByEmail(email string) (*app.User, error) {
	user, err := findUserByEmail(u.Db, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u UserService) Get(id int64) (*app.User, error) {
	//TODO implement me
	panic("implement me")
}

func findUserByEmail(db *gorm.DB, email string) (*app.User, error) {
	var user User
	result := db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return userFromModel(user), nil
}

func createUser(db *gorm.DB, user *app.User) (*app.User, error) {
	now := time.Now()
	dbUser := User{
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if result := db.Create(&dbUser); result.Error != nil {
		return nil, result.Error
	}

	return userFromModel(dbUser), nil
}

func userFromModel(user User) *app.User {
	return &app.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt.Time,
	}
}
