package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	want := Bitcoin(10)
	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
