package usecase

import (
	"financials/internal/app"
)

type ManageCurrency struct {
	CurrencyRepository app.CurrencyService
}

func NewManageCurrency(currencyRepository app.CurrencyService) *ManageCurrency {
	return &ManageCurrency{CurrencyRepository: currencyRepository}
}

type IndexCurrencyArgs struct {
	CurrencyRepository app.CurrencyService
}

func (mc ManageCurrency) Index() ([]app.Currency, error) {
	result, err := mc.CurrencyRepository.Index()
	if err != nil {
		return []app.Currency{}, err
	}

	return result, nil
}

type GetCurrencyArgs struct {
	Symbol string
}

func (mc ManageCurrency) Get(args GetCurrencyArgs) (*app.Currency, error) {
	result, err := mc.CurrencyRepository.Get(args.Symbol)
	if err != nil {
		return &app.Currency{}, err
	}

	return result, nil
}

type CreateCurrencyArgs struct {
	Symbol string `json:"symbol" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

func (mc ManageCurrency) Create(args CreateCurrencyArgs) (*app.Currency, error) {
	currency := app.NewCurrency(args.Symbol, args.Name)

	result, err := mc.CurrencyRepository.Save(currency)
	if err != nil {
		return &app.Currency{}, err
	}

	return result, nil
}

type DeleteCurrencyArgs struct {
	Symbol string
}

func (mc ManageCurrency) Delete(args DeleteCurrencyArgs) error {
	err := mc.CurrencyRepository.Delete(args.Symbol)
	if err != nil {
		return err
	}

	return nil
}
