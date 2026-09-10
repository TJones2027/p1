package customer


type Customer struct {
    name string
}


//creates cust and gives name
func NewCustomer(name string) *Customer {
    return &Customer{
        name: name,
    }
}


//enables fmt.print to use as text
func (c *Customer) String() string {
    return c.name
}
