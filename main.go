package main

import (
	"fmt"
	"p1/account"
	"p1/bank"
	"p1/customer"
)

func main() {
	myBank := bank.NewBank()
	ann := customer.NewCustomer("ann")
	bob := customer.NewCustomer("bob")
	myBank.Add(account.NewCheckingAccount(1, *ann, 100.00))
	myBank.Add(account.NewSavingAccount(2, *ann, 200.00))
	myBank.Add(account.NewSavingAccount(3, *bob, 150.00))
	myBank.Accrue(0.02)

	fmt.Println(myBank)
}