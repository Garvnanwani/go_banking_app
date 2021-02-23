package main

import (
	"banking/accounts"
	"fmt"
)

func main() {
	account := accounts.NewAccount("garv")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	account.Deposit(30)
	fmt.Println(account.Balance())
	err = account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
	fmt.Println(account)
}
