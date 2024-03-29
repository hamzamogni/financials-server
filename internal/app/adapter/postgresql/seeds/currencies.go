package seeds

import (
	"financials/internal/app/adapter/postgresql/model"
	"github.com/go-faker/faker/v4"
)

func (s Seed) CurrencySeed() {
	for i := 0; i < 10; i++ {
		currency := model.Currency{
			Name:   faker.Word(),
			Symbol: faker.Currency(),
		}

		s.db.Create(&currency)
	}
}
