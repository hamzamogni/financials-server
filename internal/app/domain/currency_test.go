package domain

import "testing"

func TestCurrency_NewCurrency(t *testing.T) {
	name := "moroccAn DiRHaM"
	symbol := "mAD"

	currency := NewCurrency(symbol, name)

	if currency.Symbol != "MAD" || currency.Name != "Moroccan Dirham" {
		t.Errorf("curency formatting failed. wanted (MAD, Moroccan Dirham), got (%s, %s)", currency.Symbol, currency.Name)
	}
}
