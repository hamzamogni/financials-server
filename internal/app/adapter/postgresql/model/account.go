package model

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	CurrencyID int
	Currency   Currency
}
