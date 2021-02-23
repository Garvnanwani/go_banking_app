package accounts

import "errors"

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
func (a Account) Balance() int {
	return a.balance
}

// Withdraw from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Cant withdraw, you are poor")
	}
	a.balance -= amount
	return nil
}
