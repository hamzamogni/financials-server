package repository

import (
	"financials/internal/app/domain"
)

type ICurrency interface {
	Index() ([]domain.Currency, error)
	Get(id string) (domain.Currency, error)
	Save(currency domain.Currency) (domain.Currency, error)
	Update(currency domain.Currency) (domain.Currency, error)
	Delete(id string) error
}
