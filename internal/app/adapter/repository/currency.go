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
			Id:     currency.ID,
			Name:   currency.Name,
			Symbol: currency.Symbol,
		})
	}

	return ret, nil
}

func (c Currency) Get(id uint) (domain.Currency, error) {
	db := postgresql.Connection()
	var currency model.Currency

	result := db.Where("id = ?", id).First(&currency)
	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return domain.Currency{
		Id:     currency.ID,
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}, nil
}

func (c Currency) Save(currency domain.Currency) (domain.Currency, error) {
	db := postgresql.Connection()

	newCurrency := model.Currency{
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}

	result := db.Create(&newCurrency)
	if result.Error != nil {
		return domain.Currency{}, result.Error
	}

	return currency, nil
}

func (c Currency) Update(currency domain.Currency) error {
	db := postgresql.Connection()

	toUpdateCurrency := model.Currency{
		ID:     currency.Id,
		Name:   currency.Name,
		Symbol: currency.Symbol,
	}

	result := db.Model(&toUpdateCurrency).Updates(toUpdateCurrency)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
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
