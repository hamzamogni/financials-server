package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
	"gorm.io/gorm"
)

type Currency struct{}

func (c Currency) Index() ([]domain.Currency, error) {
	db := postgresql.Connection()

	var ret []domain.Currency

	var currencies []model.Currency
	result := db.Find(&currencies)

	if result.Error != nil {
		return []domain.Currency{}, result.Error
	}

	for _, currency := range currencies {
		ret = append(ret, domain.Currency{
			ID:     currency.ID,
			Name:   currency.Name,
			Symbol: currency.Symbol,
		})
	}

	return ret, nil
}

func (c Currency) Get(ID string) (domain.Currency, error) {
	db := postgresql.Connection()
	var currency model.Currency

	result := db.Where("id = ?", ID).First(&currency)
	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return domain.Currency{
		ID:     currency.ID,
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil
}

func (c Currency) Save(currency domain.Currency) (domain.Currency, error) {
	db := postgresql.Connection()
	result := db.Create(&currency)

	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return currency, nil
}

func (c Currency) Update(currency domain.Currency) (domain.Currency, error) {
	var updatedCurrency model.Currency
	db := postgresql.Connection()

	if err := db.Model(&currency).Updates(currency).Error; err != nil {
		return domain.Currency{}, err
	}

	db.Where("id = ?", currency.ID).First(&updatedCurrency)
	return domain.Currency{
		ID:     updatedCurrency.ID,
		Name:   updatedCurrency.Name,
		Symbol: updatedCurrency.Symbol,
	}, nil
}

func (c Currency) Delete(ID string) error {
	var currency model.Currency
	db := postgresql.Connection()

	result := db.Delete(&currency, ID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
