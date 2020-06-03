package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 52, 435, 3, 76, 333, 3242, 13, 345, 100}
	xs := []string{"Hello", "Darko", "Marko", "Ivan", "Stjepan"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
