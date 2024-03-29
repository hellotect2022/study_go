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

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Account의 method (go 의 receiver)
func (a *Account) Deposit(amount int) {
	fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

// Withdraw x Account
func (a *Account) Withdraw(amount int) error {
	fmt.Println("Gonna withdraw", amount)
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// change of the owner
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// show the owner
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}
