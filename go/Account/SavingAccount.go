package account

import (
	"go/customer"
)

type SavingAccount struct {
	Account
}

func NewSavingAccount(number int, customer customer.Customer, balance float64) *SavingAccount {
	return &SavingAccount{
		Account: Account{
			Number:   number,
			Customer: customer,
			Balance:  balance,
		},
	}
}

func (c *SavingAccount) Accrue(rate float64) {
	c.Interest += c.Balance * rate
	c.Balance += c.Balance * rate
}
