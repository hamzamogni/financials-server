package usecase

import (
	"financials/internal/app/domain"
	"financials/internal/app/domain/repository"
)

type ManageCurrency struct {
	CurrencyRepository repository.ICurrency
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
	Id uint
}

func (mc ManageCurrency) Get(args GetCurrencyArgs) (domain.Currency, error) {
	result, err := mc.CurrencyRepository.Get(args.Id)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}

type CreateCurrencyArgs struct {
	Name   string `json:"name" binding:"required"`
	Symbol string `json:"symbol" binding:"required"`
}

func (mc ManageCurrency) Create(args CreateCurrencyArgs) (domain.Currency, error) {
	currency := domain.Currency{
		Name:   args.Name,
		Symbol: args.Symbol,
	}

	result, err := mc.CurrencyRepository.Save(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	return result, nil
}

type UpdateCurrencyArgs struct {
	Id     uint
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func (mc ManageCurrency) Update(args UpdateCurrencyArgs) (domain.Currency, error) {
	currency := domain.Currency{
		Id:     args.Id,
		Name:   args.Name,
		Symbol: args.Symbol,
	}
	err := mc.CurrencyRepository.Update(currency)
	if err != nil {
		return domain.Currency{}, err
	}

	result, err := mc.CurrencyRepository.Get(currency.Id)
	return result, nil
}

type DeleteCurrencyArgs struct {
	Id string
}

func (mc ManageCurrency) Delete(args DeleteCurrencyArgs) error {
	err := mc.CurrencyRepository.Delete(args.Id)
	if err != nil {
		return err
	}

	return nil
}
