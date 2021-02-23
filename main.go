package main

import (
	"banking/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("garv")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
