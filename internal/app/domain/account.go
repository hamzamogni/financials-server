package domain

type Account struct {
	Id       uint64
	Name     string
	Balance  float64
	Currency Currency
}
