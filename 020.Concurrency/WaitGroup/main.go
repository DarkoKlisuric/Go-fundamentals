package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS) // LINUX
	fmt.Println("ARCH\t", runtime.GOARCH) // amd64
	fmt.Println("CPUs\t", runtime.NumCPU()) // 12
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 01-simple.ChannelsBlock

	wg.Add(1) // Wait to done 01-simple.ChannelsBlock process(go routine), go foo()

	go foo()

	bar()

	fmt.Println("CPUs\t", runtime.NumCPU()) // 12
	fmt.Println("Goroutines\t", runtime.NumGoroutine()) // 2

	wg.Wait()

	/* RESULT:
	OS	 linux
	ARCH	 amd64
	CPUs	 12
	Goroutines	 01-simple.ChannelsBlock
	bar 0
	bar 01-simple.ChannelsBlock
	bar 2
	bar 3
	bar 4
	bar 5
	bar 6
	bar 7
	bar 8
	bar 9
	CPUs	 12
	Goroutines	 2
	foo 0
	foo 01-simple.ChannelsBlock
	foo 2
	foo 3
	foo 4
	foo 5
	foo 6
	foo 7
	foo 8
	foo 9
	*/
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
