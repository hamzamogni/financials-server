package usecase

import "financials/internal/app/domain/repository"

type DeleteCurrencyArgs struct {
	ID                 string
	CurrencyRepository repository.ICurrency
}

func DeleteCurrency(args DeleteCurrencyArgs) error {
	err := args.CurrencyRepository.Delete(args.ID)
	if err != nil {
		return err
	}

	return nil
}
