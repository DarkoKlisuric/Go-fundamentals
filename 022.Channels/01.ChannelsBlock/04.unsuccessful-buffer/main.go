package main

import "fmt"

func main() {
	c := make(chan int, 1) // Buffer must be 2 to run

	c <- 42
	c <- 43

	fmt.Println(<-c)
	/*
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan send]:
		main.main()
			/home/klisuric/go/src/Go-fundamentals/022.Channels/01.ChannelsBlock/04.unsuccessful-buffer/main.go:9 +0x73
		exit status 2
	*/

	/**
		With buffer 2 c := make(chan int, 2)
		fmt.Println(<-c) // 42
		fmt.Println(<-c) // 43
	 */
}
