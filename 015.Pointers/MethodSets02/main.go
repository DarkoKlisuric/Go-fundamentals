package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{"James Bond"}
	fmt.Println(p1.name)
	changeMe(&p1)
	fmt.Println(p1.name)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	// same as (*p).name = "Miss Moneypenny"
}
