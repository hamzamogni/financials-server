package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageCurrency struct {
	CurrencyRepository repository.ICurrency
}

func NewManageCurrency(currencyRepository repository.ICurrency) *ManageCurrency {
	return &ManageCurrency{CurrencyRepository: currencyRepository}
}

type IndexCurrencyArgs struct {
	CurrencyRepository repository.ICurrency
}

func (mc ManageCurrency) Index() ([]domain.Currency, error) {
	result, err := mc.CurrencyRepository.Index()
	if err != nil {
		return []domain.Currency{}, err
	}

	return result, nil
}

type GetCurrencyArgs struct {
	Symbol string
}

func (mc ManageCurrency) Get(args GetCurrencyArgs) (*domain.Currency, error) {
	result, err := mc.CurrencyRepository.Get(args.Symbol)
	if err != nil {
		return &domain.Currency{}, err
	}

	return result, nil
}

type CreateCurrencyArgs struct {
	Symbol string `json:"symbol" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

func (mc ManageCurrency) Create(args CreateCurrencyArgs) (*domain.Currency, error) {
	currency := domain.NewCurrency(args.Symbol, args.Name)

	result, err := mc.CurrencyRepository.Save(currency)
	if err != nil {
		return &domain.Currency{}, err
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
