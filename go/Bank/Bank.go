package bank

type BankAccount interface {
	account.Accruer
	String() string
}

type Bank struct {
	accounts map[BankAccount]bool
}

func NewBank() *Bank {
	return &Bank{
		accounts: make(map[BankAccount]bool),
	}
}

func (b *Bank) Add(acc BankAccount) {
	b.accounts[acc] = true
}

func (b *Bank) Accrue(rate float64) {
	for acc := range b.accounts {
		acc.Accrue(rate)
	}
}

func (b *Bank) String() string {
	str := ""
	for acc := range b.accounts {
		str += acc.String() + "\n"
	}
	return str
}
