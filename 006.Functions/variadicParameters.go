package main

import "fmt"

func main() {
	x := variadicSum(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println(x)
}

func variadicSum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T", x) // SLICES OF INT , []int

	sum := 0

	for _, value := range x {
		sum += value
	}
	fmt.Println("Sum is", sum)
	return sum
}
