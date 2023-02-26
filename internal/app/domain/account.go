package domain

import "errors"

type Account struct {
	Id       uint
	Name     string
	Balance  float64
	Currency Currency
}

func NewAccount(name string, currency Currency) *Account {
	return &Account{
		Name:     name,
		Balance:  0,
		Currency: currency,
	}
}

func (a *Account) CreditAccount(amount float64) {
	a.Balance += amount
}

func (a *Account) DebitAccount(amount float64) error {
	if amount > a.Balance {
		return errors.New("insufficient balance in account")
	}

	a.Balance -= amount
	return nil
}
