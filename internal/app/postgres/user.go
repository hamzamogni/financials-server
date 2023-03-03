package postgres

import (
	"financials/internal/app"
	"gorm.io/gorm"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService struct {
	Db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{Db: Connection()}
}

func (r UserService) Save(user *app.User) (*app.User, error) {
	newUser := User{
		Username: user.Username,
		Password: user.Password,
	}

	result := r.Db.Create(&newUser)
	if result.Error != nil {
		return &app.User{}, result.Error
	}

	return user, nil
}

func (UserService) GetByEmail(email string) (*app.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserService) Get(id int64) (*app.User, error) {
	//TODO implement me
	panic("implement me")
}
