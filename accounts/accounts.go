package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit amount to your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() string {
	return ("current balance -> " + fmt.Sprint(a.balance))
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Cant withdraw, you are poor")
	}
	a.balance -= amount
	return nil
}

// ChangeOwner changes owner of account
func (a *Account) ChangeOwner(newowner string) {
	a.owner = newowner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return a.owner + "'s account \nhas -> " + fmt.Sprint(a.balance)
}
