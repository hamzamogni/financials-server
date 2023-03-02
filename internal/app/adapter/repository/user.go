package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{Db: postgresql.Connection()}
}

func (r UserRepository) Save(user *domain.User) (*domain.User, error) {
	newUser := model.User{
		Username: user.Username,
		Password: user.Password,
	}

	result := r.Db.Create(&newUser)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}

	return user, nil
}

func (UserRepository) GetByEmail(email string) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}

func (UserRepository) Get(id int64) (*domain.User, error) {
	//TODO implement me
	panic("implement me")
}
