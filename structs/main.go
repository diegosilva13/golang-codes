package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	diego := person{
		firstName: "Diego Brener",
		lastName:  "Da Silva",
		contact: contactInfo{
			email:   "diego@email.com",
			zipCode: 78123123,
		},
	}

	diego.updateFirstName("teste")
	diego.print()
	diego.print()
}

func (p *person) updateFirstName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
