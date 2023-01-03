package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	phone int
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person

	alex.firstName = "Alexander"
	alex.lastName = "King"
	alex.contactInfo.email = "alex@ander.com"

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jim@party.com",
			phone: 1235,
		},
	}

	// fmt.Printf("%+v", alex)

	// jimPointer := &jim // memory address of jim
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)

}

func (pointerToPerson *person) updateName(newFirstName string) { // type description - says working with a pointer
	(*pointerToPerson).firstName = newFirstName // * -> value at address
}

// & - value to address
// * - address to value, but in front of type - description, work with pointers.
