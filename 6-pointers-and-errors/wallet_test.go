package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(err, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	assertNoError := func(err error) {
		t.Helper()
		if err != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(wallet, want)
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(10)
		want := Bitcoin(20)
		assertNoError(err)
		assertBalance(wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(err, ErrInsufficientFunds)
		assertBalance(wallet, Bitcoin(30))
	})
}
