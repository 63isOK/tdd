package main

// Wallet is wallet
type Wallet struct {
	balance int
}

// Deposit deposit money to bank
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance query the money
func (w *Wallet) Balance() int {
	return w.balance
}
