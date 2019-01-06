package main

import "fmt"

type contactInfo struct {
	email 	string
	zipCode int
}

type person struct {
	firstName string
	lastName	string
	contactInfo
}

func main()  {
	nested := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
	}

	nested.updateName("jimmy")
	nested.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
