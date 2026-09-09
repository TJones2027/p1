package main

import (
	"fmt"
	"p1/go/Account"
	"p1/go/Bank"
	"p1/go/customer"
)

func main() {
	myBank := bank.NewBank()
	ann := customer.NewCustomer("Ann")
	bob := customer.NewCustomer("Bob")
	myBank.Add(account.NewCheckingAccount(1, *ann, 100.00))
	myBank.Add(account.NewSavingAccount(2, *ann, 200.00))
	myBank.Add(account.NewSavingAccount(3, *bob, 150.00))
	myBank.Accrue(0.02)

	fmt.Println(myBank)
}