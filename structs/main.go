package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int 
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	everton := person{
		firstName: "Fulano",
		lastName: "Nunes",
		contactInfo: contactInfo{
			email: "test@email.com",
			zipCode: 94000000,
		},
	}

	everton.updateName("Everton")
	everton.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}