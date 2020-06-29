package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.FirstName, p.Age)
}

// ByAge implements sort.Interface for []Person based on the Age field
type ByAge []Person

type ByName []Person

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (n ByName) Len() int {
	return len(n)
}

func (n ByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n ByName) Less(i, j int) bool {
	return n[i].FirstName < n[j].FirstName
}

func main() {
	people := []Person{
		{"Darko", 25},
		{"Marko", 30},
		{"Ivan", 22},
		{"Matija", 32},
	}
	fmt.Println(people)

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
