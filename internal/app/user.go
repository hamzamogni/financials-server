package app

import "time"

type User struct {
	ID        uint
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type UserService interface {
	Save(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
	Get(id int64) (*User, error)
}
