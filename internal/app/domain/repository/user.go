package repository

import "financials/internal/app/domain"

type IUser interface {
	Save(user *domain.User) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Get(id int64) (*domain.User, error)
}
