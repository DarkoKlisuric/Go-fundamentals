package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	for index, value := range x {
		fmt.Println(index, value)
	}

	for _, value := range x {
		fmt.Println(value)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[:])
	fmt.Println(x[2:3]) // [3]

	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y)
	y = append(y, 6, 7)
	fmt.Println(y)

	z := []int{8, 9 ,10}

	y = append(y, z...)
	fmt.Println(y)
}
