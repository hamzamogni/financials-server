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
	result, err := args.CurrencyRepository.Update(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}
