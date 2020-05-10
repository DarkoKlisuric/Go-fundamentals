package main

import "fmt"

func main() {
	fmt.Println("Hello")

	func() {
		fmt.Println("World")
	}()

	func (x int) {
		fmt.Println(x)
	}(30)
}
