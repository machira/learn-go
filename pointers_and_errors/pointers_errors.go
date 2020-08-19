package pointers_and_errors

import (
	"errors"
	"fmt"
)

var insuffientFunds = errors.New("not enough money")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) (Bitcoin, error) {
	if b < w.Balance() {
		w.balance -= b
		return b, nil
	}
	return 0, insuffientFunds
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
