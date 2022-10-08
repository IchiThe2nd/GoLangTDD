package main

import (
	
	"testing"
)


func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrinsufficientFunds)
		assertBalance(t, wallet, Bitcoin(20))

	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf(" got %q want %q", got, want)
	}
}

func assertNoError  (t testing.TB, got error){
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didnt want one")

		}
}



func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didnt get one")
	}
	
	if got != want {
		t.Errorf("got %q want %q", got, want)

	}
}