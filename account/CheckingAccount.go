package account


import (
    "p1/customer"
)


type CheckingAccount struct {
    Account
}




func NewCheckingAccount(number int, customer customer.Customer, balance float64) *CheckingAccount {
    return &CheckingAccount{
        Account: Account{
            Number:   number,
            Customer: customer,
            Balance:  balance,
        },
    }
}


//cant earn accrue
func (c *CheckingAccount) Accrue(rate float64) {


}


//checks accountfor its interest
func (c *CheckingAccount) GetInterest() float64 {
    return 0
}
