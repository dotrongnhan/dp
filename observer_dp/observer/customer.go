package observer

import "fmt"

type Customer struct {
	ID string
}

func (c *Customer) Update(str string) {
	fmt.Printf("Sending mail to customer %s for item %s \n", c.ID, str)
}

func (c *Customer) GetId() string {
	return c.ID
}
