package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageAccount struct {
	AccountRepository  repository.IAccount
	CurrencyRepository repository.ICurrency
}

func NewManageAccount(ar repository.IAccount, cr repository.ICurrency) *ManageAccount {
	return &ManageAccount{AccountRepository: ar, CurrencyRepository: cr}
}

type IndexAccountArgs struct {
	AccountRepository repository.IAccount
}

func (ma ManageAccount) Index() ([]domain.Account, error) {
	result, err := ma.AccountRepository.Index()
	if err != nil {
		return []domain.Account{}, err
	}

	return result, nil
}

type GetAccountArgs struct {
	Id uint
}

func (ma ManageAccount) Get(args GetAccountArgs) (*domain.Account, error) {
	result, err := ma.AccountRepository.Get(args.Id)
	if err != nil {
		return &domain.Account{}, err
	}

	return result, nil
}

type CreateAccountArgs struct {
	Name           string `json:"name" binding:"required"`
	CurrencySymbol string `json:"currency_symbol" binding:"required"`
}

func (ma ManageAccount) Create(args CreateAccountArgs) (*domain.Account, error) {
	currency, err := ma.CurrencyRepository.Get(args.CurrencySymbol)
	if err != nil {
		return &domain.Account{}, err
	}

	account := &domain.Account{
		Name:     args.Name,
		Balance:  0,
		Currency: *currency,
	}

	newAccount, err := ma.AccountRepository.Save(account)
	if err != nil {
		return &domain.Account{}, err
	}

	return newAccount, nil
}
