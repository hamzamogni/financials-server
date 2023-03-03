package postgres

import (
	"financials/internal/app"
	"gorm.io/gorm"
)

type Currency struct {
	Symbol string `gorm:"primarykey"`
	Name   string
}

type CurrencyService struct {
	Db *gorm.DB
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		Db: Connection(),
	}
}

func (cr CurrencyService) Index() ([]app.Currency, error) {
	var ret []app.Currency

	var currencies []Currency
	result := cr.Db.Find(&currencies)

	if result.Error != nil {
		return []app.Currency{}, result.Error
	}

	for _, currency := range currencies {
		ret = append(ret, app.Currency{
			Name:   currency.Name,
			Symbol: currency.Symbol,
		})
	}

	return ret, nil
}

func (cr CurrencyService) Get(symbol string) (*app.Currency, error) {
	var currency Currency

	result := cr.Db.Where("symbol = ?", symbol).First(&currency)
	if result.Error != nil {
		return &app.Currency{}, result.Error
	}

	return &app.Currency{
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil
}

func (cr CurrencyService) Save(currency *app.Currency) (*app.Currency, error) {
	newCurrency := Currency{
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}

	result := cr.Db.Create(&newCurrency)
	if result.Error != nil {
		return &app.Currency{}, result.Error
	}

	return currency, nil
}

func (cr CurrencyService) Delete(symbol string) error {
	result := cr.Db.Model(&Currency{}).Where("symbol = ?", symbol).Delete(&Currency{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
