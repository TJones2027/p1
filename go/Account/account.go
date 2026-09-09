package account

import (
	"fmt"
	"go/customer"
)

type Account struct {
	Number   int
	Balance  float64
	Customer customer.Customer
}

type Accruer interface {
	Accrue(rate float64)
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) {
	a.Balance -= amount
}

func (a *Account) String() string {
	return fmt.Sprintf("%d: %s: %f", a.Number, a.Customer, a.Balance)
}
