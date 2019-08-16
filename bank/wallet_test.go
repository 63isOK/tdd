package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	})
}
