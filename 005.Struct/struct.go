package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}
func main() {
	p1 := Person{
		firstName: "Darko",
		lastName: "Klisurić",
		}
	p2 := Person{
		firstName: "Marko",
		lastName:  "Markić",
	}
	fmt.Println(p1)
	fmt.Println(p1.lastName)
	fmt.Println(p2)
	fmt.Println(p2.lastName)

}