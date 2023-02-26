package model

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	Balance    float64
	CurrencyID string
	Currency   Currency
}
