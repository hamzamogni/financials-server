package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type GetCurrencyArgs struct {
	ID                 string
	CurrencyRepository repository.ICurrency
}

func GetCurrency(args GetCurrencyArgs) (domain.Currency, error) {
	result, err := args.CurrencyRepository.Get(args.ID)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}
