package app

type User struct {
	Username string
	Email    string
	Password string
}

type UserService interface {
	Save(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
	Get(id int64) (*User, error)
}

type AuthService interface {
	SignUp(user *User) *User
}
