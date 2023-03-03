package seed

import (
	"financials/internal/app/postgres"
	"github.com/go-faker/faker/v4"
)

func (s Seed) CurrencySeed() {
	for i := 0; i < 10; i++ {
		currency := postgres.Currency{
			Name:   faker.Word(),
			Symbol: faker.Currency(),
		}

		s.db.Create(&currency)
	}
}
