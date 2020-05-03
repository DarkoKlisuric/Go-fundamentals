package main

import "fmt"

func main() {
	var x [5]int
	x[3] = 43
	fmt.Println(x)

	var y = []int{1, 2, 3, 4}

	fmt.Println(y)
}
