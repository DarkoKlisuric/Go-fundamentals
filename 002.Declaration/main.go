package main

import "fmt"

var b = 100
// declare there is a variable with the identifier "c"
// false for booleans, 0 for numeric types, "" for strings,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var c int

func main()  {
	// DECLARE a variable and assign a value (of a certain type)
	a := 100
	fmt.Println(a)
	a = 200
	fmt.Println(a)

	fmt.Println(b)
	b = 200
	fmt.Println(b)

	foo()
}

func foo()  {
	fmt.Println("Again;", b)
}
