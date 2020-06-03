package main

import "fmt"

func main() {
	c := make(chan <- int, 2) // Now can only send a value into channel
	c2 := make(<- chan int, 2)
	c <- 42
	c <- 43

	c2 <- 42 // Receive only!
	c2 <- 43 //

	fmt.Println(<-c) // Can't read
	fmt.Println(<-c)

	fmt.Println(<-c2) // can read
	fmt.Println(<-c2)
}
