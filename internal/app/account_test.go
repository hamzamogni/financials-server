package app

import "testing"

func TestAccount_CreditAccount(t *testing.T) {
	acc := NewAccount("CFG", Currency{})

	addedAmount := 120.5
	acc.CreditAccount(addedAmount)

	got := acc.Balance

	if got != addedAmount {
		t.Errorf("got %f, wanted %f", got, addedAmount)
	}
}
