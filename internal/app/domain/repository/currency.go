package repository

import (
	"financials/internal/app/domain"
)

type ICurrency interface {
	Index() ([]domain.Currency, error)
	Get(symbol string) (*domain.Currency, error)
	Save(currency *domain.Currency) (*domain.Currency, error)
	Delete(symbol string) error
}
