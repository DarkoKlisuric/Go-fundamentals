package main

import "fmt"

func main() {
	// for loop

	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("The outher loop: %d\t The inner loop %d\t Result %d\n", i, j, i*j)
		}
	}

	// Standard while loop
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("Done")

	// Infinite loop - for(;;)
	y := 5
	for {
		if y > 25 {
			fmt.Println("Done")
			break
		}
		y++
	}
	// Printig ASCII
	for k := 33; k <= 122; k++ {
		fmt.Printf("%v\t%#U\n", k, k)
	}

}
