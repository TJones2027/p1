package account


import (
    "fmt"
    "p1/customer"
)






type Account struct {
    Number   int
    Balance  float64
    Customer customer.Customer
}




type Accruer interface {
    Accrue(rate float64)
}


//function to get account balance
func (a *Account) GetBalance() float64 {
    return a.Balance
}


//adds Deposit function
func (a *Account) Deposit(amount float64) {
    //takes balance and adds amount
    a.Balance += amount
}


//adds the withdraw function
func (a *Account) Withdraw(amount float64) {
    //takes balance and subtracts amount
    a.Balance -= amount
}


func (a *Account) String() string {
    return fmt.Sprintf("%d: %s: %.2f", a.Number, &a.Customer, a.Balance)
}
