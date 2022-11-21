package repository

import "goland/internal/app/domain"

type ICurrency interface {
	Get(ID string) (domain.Currency, error)
	Save(currency domain.Currency) (domain.Currency, error)
}
