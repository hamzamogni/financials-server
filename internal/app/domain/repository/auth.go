package repository

import "financials/internal/app/domain"

type Auth interface {
	SignUp(user *domain.User) *domain.User
}
