package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type GetCurrencyArgs struct {
	Id                 uint
	CurrencyRepository repository.ICurrency
}

func GetCurrency(args GetCurrencyArgs) (domain.Currency, error) {
	result, err := args.CurrencyRepository.Get(args.Id)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}
