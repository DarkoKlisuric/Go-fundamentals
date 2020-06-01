package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
}

func (p *person) speak() {
	fmt.Println("Hello")
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

	saySomething(&darko) // Hello
	darko.speak() // Hello
}
