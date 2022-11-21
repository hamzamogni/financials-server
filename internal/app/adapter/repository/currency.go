package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
)

type Currency struct{}

func (c Currency) Get(ID string) (domain.Currency, error) {
	db := postgresql.Connection()
	var currency model.Currency

	result := db.Where("id = ?", ID).First(&currency)
	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return domain.Currency{
		ID:     currency.ID,
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil
}

func (c Currency) Save(currency domain.Currency) (domain.Currency, error) {
	db := postgresql.Connection()
	result := db.Create(&currency)

	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return currency, nil
}
