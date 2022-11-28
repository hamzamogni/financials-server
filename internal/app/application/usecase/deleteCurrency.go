package usecase

import "financials/internal/app/domain/repository"

type DeleteCurrencyArgs struct {
	Id                 string
	CurrencyRepository repository.ICurrency
}

func DeleteCurrency(args DeleteCurrencyArgs) error {
	err := args.CurrencyRepository.Delete(args.Id)
	if err != nil {
		return err
	}

	return nil
}
