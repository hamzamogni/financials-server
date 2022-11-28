package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type UpdateCurrencyArgs struct {
	Id                 uint
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	CurrencyRepository repository.ICurrency
}

func UpdateCurrency(args UpdateCurrencyArgs) (domain.Currency, error) {
	currency := domain.Currency{
		Id:     args.Id,
		Name:   args.Name,
		Symbol: args.Symbol,
	}
	err := args.CurrencyRepository.Update(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	result, err := args.CurrencyRepository.Get(currency.Id)
	return result, nil
}
