package main

import "fmt"

func main() {
	a := 50
	fmt.Println(a)  // 50
	fmt.Println(&a) // 0xc4200180e0 gives you the address

	fmt.Printf("%T\n", a)  // int
	fmt.Printf("%T\n", &a) // *int

	var b *int = &a
	fmt.Println(b)  // 0xc4200180e0
	fmt.Println(*b) // 50 - gives you the value stored at an anddress when you have the address

	*b = 100
	fmt.Println(b) // 0xc4200180e0
	fmt.Println(a) // 100

	x := 0
	foo(x)
	fmt.Println(x) // 0

	y := 0
	bar(&y)
	fmt.Println(y) // 43
}

func foo(x int) {
	fmt.Println(x) // 0
	x = 43 
	fmt.Println(x) // 43
}

func bar(x *int) {
	fmt.Println(*x) // 0
	*x = 43
	fmt.Println(*x) // 43
}
