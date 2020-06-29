package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type programer struct {
	person
	working bool
}

func main() {
	p1 := person{
		firstName: "Darko",
		lastName:  "Klisurić",
	}
	p2 := person{
		firstName: "Marko",
		lastName:  "Markić",
	}

	pro1 := programer{
		person: person{
			firstName: "Darko",
			lastName:  "Klisueic",
		},
		working: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.lastName)
	fmt.Println(p2)
	fmt.Println(p2.lastName)
	fmt.Println(pro1, pro1.person.lastName, pro1.person.firstName, pro1.working)

	p3 := struct {
		firstName string
		lastName  string
	}{
		firstName: "Darko",
		lastName:  "Klisuric",
	}

	fmt.Println(p3)
}
