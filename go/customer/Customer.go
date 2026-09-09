package customer

type customer struct {
	name string
}

func NewCustomer(name string) *customer {
	return &customer{
		name: name,
	}
}

func (c *customer) String() string {
	return c.name
}
