package main

import "fmt"

func main() {
	fmt.Println(factorialRecursion(5))
	fmt.Println(factorialLoop(5))
}

func factorialRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursion(n-1)
}

func factorialLoop(n int) int {
	total := 1

	for ; n > 0; n-- {
		total *= n
	}

	return total
}
