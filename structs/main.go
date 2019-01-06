package main

import "fmt"

type person struct {
	firstName string
	lastName	string
}

func main()  {
	alex := person{
		firstName:"Alex",
		lastName: "Anderson",
	}

	var robert person
	robert.firstName = "Robert"

	fmt.Println(alex, robert)
	fmt.Printf("%+v", alex)
}
