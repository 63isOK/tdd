package main

import (
	"errors"
	"fmt"
)

var (
	withdrawError = errors.New("amount error")
)

// Bitcoin is bit coin
type Bitcoin uint

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit deposit money to bank
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw withdraw money from bank
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount <= w.balance {
		w.balance -= amount
		return nil
	}

	// return nil
	return withdrawError
}

// Balance query the money
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{Bitcoin(20)}
	err := w.Withdraw(100)
	if err != nil {
		fmt.Println(w.Balance())
	}
}
