package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type IndexCurrencyArgs struct {
	CurrencyRepository repository.ICurrency
}

func IndexCurrency(args IndexCurrencyArgs) ([]domain.Currency, error) {
	result, err := args.CurrencyRepository.Index()
	if err != nil {
		return []domain.Currency{}, err
	}

	return result, nil
}
