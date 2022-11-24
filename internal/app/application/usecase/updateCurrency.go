package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type UpdateCurrencyArgs struct {
	ID                 uint
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	CurrencyRepository repository.ICurrency
}

func UpdateCurrency(args UpdateCurrencyArgs) (domain.Currency, error) {
	currency := domain.Currency{
		ID:     args.ID,
		Name:   args.Name,
		Symbol: args.Symbol,
	}
	err := args.CurrencyRepository.Update(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	result, err := args.CurrencyRepository.Get(currency.ID)
	return result, nil
}
