package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageCurrency struct{}

type IndexCurrencyArgs struct {
	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Index(args IndexCurrencyArgs) ([]domain.Currency, error) {
	result, err := args.CurrencyRepository.Index()
	if err != nil {
		return []domain.Currency{}, err
	}

	return result, nil
}

type GetCurrencyArgs struct {
	Id                 uint
	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Get(args GetCurrencyArgs) (domain.Currency, error) {
	result, err := args.CurrencyRepository.Get(args.Id)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}

type CreateCurrencyArgs struct {
	Name   string `json:"name" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`

	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Create(args CreateCurrencyArgs) (domain.Currency, error) {
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

type UpdateCurrencyArgs struct {
	Id                 uint
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Update(args UpdateCurrencyArgs) (domain.Currency, error) {
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

type DeleteCurrencyArgs struct {
	Id                 string
	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Delete(args DeleteCurrencyArgs) error {
	err := args.CurrencyRepository.Delete(args.Id)
	if err != nil {
		return err
	}

	return nil
}
