package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type CreateCurrencyArgs struct {
	Name   string `json:"name" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`

	CurrencyRepository repository.ICurrency
}

func CreateCurrency(args CreateCurrencyArgs) (domain.Currency, error) {
	currency := domain.Currency{
		Name:   args.Name,
		Symbol: args.Symbol,
	}

	result, err := args.CurrencyRepository.Save(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}
