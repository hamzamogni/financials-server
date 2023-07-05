package usecase

import (
	"financials/internal/app"
)

type ManageUser struct {
	UserService app.UserService
}

func NewManageUser(userService app.UserService) *ManageUser {
	return &ManageUser{UserService: userService}
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

	newUser, err := mu.UserService.Save(user)
	if err != nil {
		return &app.User{}, err
	}

	return newUser, nil

}
