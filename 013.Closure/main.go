package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a()) // 01.ChannelsBlock
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(b()) // 01.ChannelsBlock
	fmt.Println(b()) // 2
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
