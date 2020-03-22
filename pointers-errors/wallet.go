package main

import (
	"errors"
	"fmt"
)

// Bitcoin defination
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet model
type Wallet struct {
	balance Bitcoin
}

// Balance shows how much is in the Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit adds an amount to the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds - Error message to return when there is insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw removes an amount from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
