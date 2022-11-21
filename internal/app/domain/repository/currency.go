package repository

import "financials/internal/app/domain"

type ICurrency interface {
	Get(ID string) (domain.Currency, error)
	Save(currency domain.Currency) (domain.Currency, error)
}
