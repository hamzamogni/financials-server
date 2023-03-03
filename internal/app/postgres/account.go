package postgres

import (
	"financials/internal/app"
	"fmt"
	"gorm.io/gorm"
)

type Account struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"uniqueIndex"`
	Balance    float64
	CurrencyID string
	Currency   Currency
}

type AccountService struct {
	Db *gorm.DB
}

func NewAccountService() *AccountService {
	return &AccountService{Db: Connection()}
}

func (ar AccountService) Index() ([]app.Account, error) {
	var ret []app.Account
	var accounts []Account
	results := ar.Db.Joins("Currency").Find(&accounts)
	fmt.Println(accounts)

	if results.Error != nil {
		return []app.Account{}, results.Error
	}

	for _, account := range accounts {
		ret = append(ret, app.Account{
			Id:      account.ID,
			Name:    account.Name,
			Balance: account.Balance,
			Currency: app.Currency{
				Symbol: account.Currency.Symbol,
				Name:   account.Currency.Name,
			},
		})
	}

	return ret, nil
}

func (ar AccountService) Get(id uint) (*app.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (ar AccountService) Update(account *app.Account) error {
	//TODO implement me
	panic("implement me")
}

func (ar AccountService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (ar AccountService) Save(account *app.Account) (*app.Account, error) {
	newAccount := Account{
		Name: account.Name,
		Currency: Currency{
			Symbol: account.Currency.Symbol,
		},
	}

	if err := ar.Db.Create(&newAccount).Error; err != nil {
		return &app.Account{}, err
	}

	return account, nil
}
