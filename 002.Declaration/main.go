package main

import (
	"fmt"
)

var b = 100

// declare there is a variable with the identifier "c"
// false for booleans, 0 for numeric types, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var c int

type Darko string

var d Darko

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	// DECLARE a variable and assign a value (of a certain type)
	a := 100
	fmt.Println(a)
	a = 200
	fmt.Println(a)

	fmt.Println(b)
	b = 200
	fmt.Println(b)

	foo()

	fmt.Printf("%T\n", d) // main.Darko type

	iotaDeclaration()
}

func foo() {
	fmt.Println("Again;", b)
}

func iotaDeclaration() {
	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
