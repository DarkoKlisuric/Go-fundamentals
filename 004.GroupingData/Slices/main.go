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

	fmt.Println(x) // [ 01.ChannelsBlock, 2, 3, 4, 5 ]
	fmt.Println(len(x)) // 5
	fmt.Println(x[:]) // [01.ChannelsBlock, 2, 3, 4, 5]
	fmt.Println(x[2:3]) // [3]

	y := []int{1, 2, 3, 4, 5}
	fmt.Println(y) // [01.ChannelsBlock, 2, 3, 4, 5]
	y = append(y, 6, 7)
	fmt.Println(y) // [01.ChannelsBlock, 2, 3, 4, 5 6, 7]

	z := []int{8, 9 ,10}

	y = append(y, z...)
	fmt.Println(y) // [01.ChannelsBlock, 2, 3, 4, 5, 6, 7, 8 , 9, 10]

	y = append(y[:2], y[4:]...)

	fmt.Println(y) //  [01.ChannelsBlock, 2, 5, 6, 7, 8 , 9, 10]

	k := make([]int, 10, 100)

	fmt.Println(k)
	fmt.Println(len(k))
	fmt.Println(cap(k))

	dk := []string{"Darko", "Klisuric"}
	mk := []string{"Marko", "Markic"}

	xp := [][]string{dk, mk}

	fmt.Println(xp)
}
