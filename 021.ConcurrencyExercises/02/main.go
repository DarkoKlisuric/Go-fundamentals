package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
}

type person2 struct {
	FirstName string
	LastName  string
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func (p2 *person2) speak() {
	fmt.Println("Holla")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	darko := person{
		FirstName: "Darko",
		LastName:  "Klisuric",
	}

	darko2 := person2{
		FirstName: "Darko",
		LastName:  "Klisuric",
	}
	saySomething(&darko)  // Hello
	saySomething(&darko2) // Holla
	darko.speak()         // Hello
	darko2.speak()        // Holla
}
