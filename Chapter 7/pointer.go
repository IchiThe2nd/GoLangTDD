package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf(" balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

var ErrinsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrinsufficientFunds
	}

	w.balance -= amount
	return nil

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
