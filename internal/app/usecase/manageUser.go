package usecase

import (
	"financials/internal/app"
)

type ManageUser struct {
	UserRepository app.UserService
}

func NewManageUser(userRepository app.UserService) *ManageUser {
	return &ManageUser{UserRepository: userRepository}
}

type CreateUserArgs struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (mu ManageUser) Create(args CreateUserArgs) (*app.User, error) {
	user := &app.User{
		Username: args.Username,
		Email:    args.Email,
		Password: args.Password,
	}

	newUser, err := mu.UserRepository.Save(user)
	if err != nil {
		return &app.User{}, err
	}

	return newUser, nil

}
