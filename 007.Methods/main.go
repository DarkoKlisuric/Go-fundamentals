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

func (p programer) programming() {
	fmt.Println("Hello from method")
}

func main() {
	p1 := programer{
		person: person{
			firstName: "Darko",
			lastName:  "Klisuric",
		},
		working: true,
	}
	fmt.Println(p1)
	p1.programming()
}
