package usecase

import (
	"financials/internal/app"
)

type ManageAccount struct {
	AccountRepository  app.AccountService
	CurrencyRepository app.CurrencyService
}

func NewManageAccount(ar app.AccountService, cr app.CurrencyService) *ManageAccount {
	return &ManageAccount{AccountRepository: ar, CurrencyRepository: cr}
}

type IndexAccountArgs struct {
	AccountRepository app.AccountService
}

func (ma ManageAccount) Index() ([]app.Account, error) {
	result, err := ma.AccountRepository.Index()
	if err != nil {
		return []app.Account{}, err
	}

	return result, nil
}

type GetAccountArgs struct {
	Id uint
}

func (ma ManageAccount) Get(args GetAccountArgs) (*app.Account, error) {
	result, err := ma.AccountRepository.Get(args.Id)
	if err != nil {
		return &app.Account{}, err
	}

	return result, nil
}

type CreateAccountArgs struct {
	Name           string `json:"name" binding:"required"`
	CurrencySymbol string `json:"currency_symbol" binding:"required"`
}

func (ma ManageAccount) Create(args CreateAccountArgs) (*app.Account, error) {
	currency, err := ma.CurrencyRepository.Get(args.CurrencySymbol)
	if err != nil {
		return &app.Account{}, err
	}

	account := &app.Account{
		Name:     args.Name,
		Balance:  0,
		Currency: *currency,
	}

	newAccount, err := ma.AccountRepository.Save(account)
	if err != nil {
		return &app.Account{}, err
	}

	return newAccount, nil
}
