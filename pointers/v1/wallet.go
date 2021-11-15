package wallet

import "fmt"

type (
	Wallet struct {
		balance Bitcoin
	}
	Bitcoin  int
	Stringer interface {
		String() string
	}
)

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
