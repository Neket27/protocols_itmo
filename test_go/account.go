package test_go

import "math/rand"

type Account struct {
	id      int
	balance float64
	owner   string
}

func NewAccountWithDetails(id int, balance float64, owner string) *Account {
	return &Account{id, balance, owner}
}

func NewAccount(owner string) *Account {
	return &Account{rand.Intn(10), 0, owner}
}

func (a *Account) Deposit(amount float64) {
	a.balance += amount
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a Account) Deposit2(amount float64) *Account {
	a.balance += amount
	return &a
}

func (a Account) GetBalance2() float64 {
	return a.balance
}
