package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
	"gorm.io/gorm"
)

type CurrencyRepository struct {
	Db *gorm.DB
}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{
		Db: postgresql.Connection(),
	}
}

func (cr CurrencyRepository) Index() ([]domain.Currency, error) {
	var ret []domain.Currency

	var currencies []model.Currency
	result := cr.Db.Find(&currencies)

	if result.Error != nil {
		return []domain.Currency{}, result.Error
	}

	for _, currency := range currencies {
		ret = append(ret, domain.Currency{
			Name:   currency.Name,
			Symbol: currency.Symbol,
		})
	}

	return ret, nil
}

func (cr CurrencyRepository) Get(symbol string) (*domain.Currency, error) {
	var currency model.Currency

	result := cr.Db.Where("symbol = ?", symbol).First(&currency)
	if result.Error != nil {
		return &domain.Currency{}, result.Error
	}

	return &domain.Currency{
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil
}

func (cr CurrencyRepository) Save(currency *domain.Currency) (*domain.Currency, error) {
	newCurrency := model.Currency{
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}

	result := cr.Db.Create(&newCurrency)
	if result.Error != nil {
		return &domain.Currency{}, result.Error
	}

	return currency, nil
}

func (cr CurrencyRepository) Delete(symbol string) error {
	result := cr.Db.Model(&model.Currency{}).Where("symbol = ?", symbol).Delete(&model.Currency{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
