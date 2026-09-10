package main


import (
    "fmt"
    "p1/account"
    "p1/bank"
    "p1/customer"
)


func main() {
    myBank := bank.NewBank()
    //adds name for cust
    ann := customer.NewCustomer("ann")
    bob := customer.NewCustomer("bob")
    //adds account calling from Checking and Savings
    myBank.Add(account.NewCheckingAccount(1, *ann, 100.00))
    myBank.Add(account.NewSavingAccount(2, *ann, 200.00))
    myBank.Add(account.NewSavingAccount(3, *bob, 150.00))
    myBank.Accrue(0.02)


    //calls Bank.String() to produce final output
    fmt.Println(myBank)
}
