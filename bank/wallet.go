package main

import (
	"errors"
	"fmt"
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

	return errors.New("amount error")
}

// Balance query the money
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{Bitcoin(20)}
	w.Withdraw(100)
	fmt.Println(w.Balance())
}
