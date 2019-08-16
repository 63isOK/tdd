package main

import (
	"fmt"
)

// Bitcoin is bit coin
type Bitcoin int

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

// Balance query the money
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {
	w := Wallet{}
	w.Deposit(1)
	fmt.Println(w.Balance())
}
