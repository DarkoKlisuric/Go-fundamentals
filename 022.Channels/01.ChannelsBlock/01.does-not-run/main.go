package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)

	/*
	fatal error: all goroutines are asleep - deadlock!

	goroutine 1 [chan send]:
	main.main()
		/home/klisuric/go/src/Go-fundamentals/022.Channels/01.ChannelsBlock/01.does-not-run/main.go:8 +0x59
	exit status 2
	*/
}

