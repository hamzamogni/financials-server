package repository

import (
	"financials/internal/app/adapter/postgresql"
	"financials/internal/app/adapter/postgresql/model"
	"financials/internal/app/domain"
	"fmt"
	"gorm.io/gorm"
)

type AccountRepository struct {
	Db *gorm.DB
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{Db: postgresql.Connection()}
}

func (ar AccountRepository) Index() ([]domain.Account, error) {
	var ret []domain.Account
	var accounts []model.Account
	results := ar.Db.Joins("Currency").Find(&accounts)
	fmt.Println(accounts)

	if results.Error != nil {
		return []domain.Account{}, results.Error
	}

	for _, account := range accounts {
		ret = append(ret, domain.Account{
			Id:      account.ID,
			Name:    account.Name,
			Balance: account.Balance,
			Currency: domain.Currency{
				Symbol: account.Currency.Symbol,
				Name:   account.Currency.Name,
			},
		})
	}

	return ret, nil
}

func (ar AccountRepository) Get(id uint) (*domain.Account, error) {
	//TODO implement me
	panic("implement me")
}

func (ar AccountRepository) Update(account *domain.Account) error {
	//TODO implement me
	panic("implement me")
}

func (ar AccountRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (ar AccountRepository) Save(account *domain.Account) (*domain.Account, error) {
	newAccount := model.Account{
		Name: account.Name,
		Currency: model.Currency{
			Symbol: account.Currency.Symbol,
		},
	}

	if err := ar.Db.Create(&newAccount).Error; err != nil {
		return &domain.Account{}, err
	}

	return account, nil
}
