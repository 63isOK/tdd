package main

// Wallet is wallet
type Wallet struct{}

// Deposit deposit money to bank
func (w Wallet) Deposit(amount int) {
}

// Balance query the money
func (w Wallet) Balance() int {
	return 0
}
