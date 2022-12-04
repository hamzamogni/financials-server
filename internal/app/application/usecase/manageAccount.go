package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageAccount struct {
	AccountRepository  repository.IAccount
	CurrencyRepository repository.ICurrency
}

type CreateAccountArgs struct {
	Name       string `json:"name" binding:"required"`
	CurrencyId uint   `json:"currency_id" binding:"required"`
}

func (ma ManageAccount) Create(args CreateAccountArgs) (domain.Account, error) {
	currency, err := ma.CurrencyRepository.Get(args.CurrencyId)
	if err != nil {
		return domain.Account{}, err
	}

	account := domain.Account{
		Name:     args.Name,
		Balance:  0,
		Currency: currency,
	}

	newAccount, err := ma.AccountRepository.Save(account)
	if err != nil {
		return domain.Account{}, err
	}

	return newAccount, nil
}
