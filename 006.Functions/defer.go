package main

import "fmt"

func main() {
	defer x()
	y := y()
	fmt.Println(y)
}

func x()  {
	fmt.Println("x")
}

func y() string {
	return fmt.Sprintln("y")
}
