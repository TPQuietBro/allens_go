package bank

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name string
	Address string
	Phone string
}

type Account struct {
	Customer
	Number int32
	Balance float64
}

func (a *Account)Deposit(amount float64) error  {
	if amount <= 0 {
		return errors.New("mount is invalid")
	}
	a.Balance += amount
	fmt.Println("balance = ", a.Balance)
	return nil
}