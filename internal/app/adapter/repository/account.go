package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
)

type Account struct {
}

func (a Account) Save(account domain.Account) (domain.Account, error) {
	db := postgresql.Connection()

	newAccount := model.Account{
		Name: account.Name,
		Currency: model.Currency{
			ID: account.Currency.Id,
		},
	}

	if err := db.Create(&newAccount).Error; err != nil {
		return domain.Account{}, err
	}

	return account, nil
}
