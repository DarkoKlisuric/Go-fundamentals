package main

import "fmt"

func main() {
	foo()
	bar("Darko")
	s1 := woo("Moneypeeny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// func (r reciver) identifier(parameters) (return(s)) { ... }
func foo() {
	fmt.Println("Hello from foo")
}

// EVERYTHING IN GO IS PASS BY VALUE
func bar(s string)  {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from wo,",s)
}

func mouse(firstName string, lastName string) (string, bool) {
	a := fmt.Sprint(firstName, lastName, `Say "Hello"`)
	b := false
	return a, b
}