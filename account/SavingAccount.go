package account


import (
    "p1/customer"
)


type SavingAccount struct {
    Account
    Interest float64
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


//calculates accrue
func (c *SavingAccount) Accrue(rate float64) {
    c.Interest += c.Balance * rate
    c.Balance += c.Balance * rate
}
//grabs interest for Savings
func (c *SavingAccount) GetInterest() float64 {
    return c.Interest
}
