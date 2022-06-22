package bank

import "testing"

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:"john",
			Address:"lac",
			Phone:"1212121",
		},
		Number: 1101,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error("account is null")
	}
}