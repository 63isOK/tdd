package main

import "testing"

func TestWallet(t *testing.T) {

	assertWallet := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s got %s", want, got)
		}
	}

	assertError := func(t *testing.T, got, want error) {
		t.Helper()

		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertWallet(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertWallet(t, wallet, Bitcoin(10))

		if err != nil {
			t.Errorf("shoudn't have a error")
		}
	})

	t.Run("illegal operation", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(100)
		assertWallet(t, wallet, Bitcoin(20))
		assertError(t, err, withdrawError)
	})
}
