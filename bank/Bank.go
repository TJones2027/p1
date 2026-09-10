package bank


import (
    "fmt"
    "p1/account"
)


//interface any account
type BankAccount interface {
    account.Accruer
    String() string
    GetInterest() float64
}


//hold set of accounts
type Bank struct {
    accounts map[BankAccount]bool
}


//builds new bank
func NewBank() *Bank {
    return &Bank{
        accounts: make(map[BankAccount]bool),
    }
}


//ads account to bank
func (b *Bank) Add(acc BankAccount) {
    b.accounts[acc] = true
}




//applies rate to every account
func (b *Bank) Accrue(rate float64) {
    for acc := range b.accounts {
        acc.Accrue(rate)
    }
}


// gets interest across all accounts w/ goroutines
func (b *Bank) TotalInterest() float64 {
    results := make(chan float64, len(b.accounts))


    for acc := range b.accounts {
        go func(a BankAccount) {
            results <- a.GetInterest()
        }(acc)
    }


    total := 0.0
    for i := 0; i < len(b.accounts); i++ {
        total += <-results // blocks until a result is ready
    }
    return total
}


// toString() equivalent - used automatically by fmt.Println.
func (b *Bank) String() string {
    str := ""
    for acc := range b.accounts {
        str += acc.String() + "\n"
    }
    str += fmt.Sprintf("Total Interest: %.2f", b.TotalInterest())
    return str
}
