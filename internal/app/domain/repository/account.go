package repository

import "financials/internal/app/domain"

type IAccount interface {
	Index() ([]domain.Account, error)
	Get(id uint) (*domain.Account, error)
	Save(account *domain.Account) (*domain.Account, error)
	Update(account *domain.Account) error
	Delete(id string) error
}
