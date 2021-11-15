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
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
