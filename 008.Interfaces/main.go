package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}
type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("Hello from method", s.firstName, s.lastName, "secret agent speak")
}

func (p person) speak() {
	fmt.Println("Hello from method", p.firstName, p.lastName, "person speak")
}

type human interface {
	speak()
}

func bar(h human)  {
	fmt.Println("I was passed in bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
		},
		ltk: true,
	}

	p1 := person{
		firstName: "Doctor",
		lastName: "No",
	}

	fmt.Println(p1) // {Doctor No}

	sa1.speak() // Hello from method James Bond secret agent speak

	bar(p1) // I was passed in bar {Doctor No}

	bar(sa1) // I was passed in bar {{James Bond} true}
}
