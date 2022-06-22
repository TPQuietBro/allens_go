package main

import (
	"fmt"
	"net/http"

	"github.com/msft/bank"
)

var accounts = map[float64]*bank.Account{}

func main()  {
	account := bank.Account{
		Customer: bank.Customer{
			Name: "allen",
			Address: "lac",
			Phone: "136",
		},
		Number: 001,
		Balance: 1000,
	}
	fmt.Println(account.Deposit(10));
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberReq := req.URL.Query().Get("number")
	if numberReq == "" {
		fmt.Println(" number is nil")
		return
	}

	if  number, err := strconv.ParseFloag(numberReq, 64); err != nil {
		fmt.Println("number is invalid: " ,number);
	} else {
		account, ok := accounts[numberReq]
		if !ok {
			fmt.Println("not found account")
		} else {
			fmt.Fprintf(w, account.Statement())
		}
	}
}