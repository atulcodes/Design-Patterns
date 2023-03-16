package main

import "fmt"

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Println("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getId() string {
	if c == nil {
		return ""
	}
	return c.id
}
