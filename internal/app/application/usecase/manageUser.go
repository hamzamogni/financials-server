package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageUser struct {
	UserRepository repository.IUser
}

func NewManageUser(userRepository repository.IUser) *ManageUser {
	return &ManageUser{UserRepository: userRepository}
}

type CreateUserArgs struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (mu ManageUser) Create(args CreateUserArgs) (*domain.User, error) {
	user := &domain.User{
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	}

	newUser, err := mu.UserRepository.Save(user)
	if err != nil {
		return &domain.User{}, err
	}

	return newUser, nil

}
