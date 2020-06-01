package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println(runtime.NumGoroutine()) // 1

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hello from one")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine()) // 2

	go func() {
		fmt.Println("Hello from two")
		wg.Done()
	}()
	fmt.Println(runtime.NumGoroutine()) // 3

	wg.Wait()

	fmt.Println(runtime.NumGoroutine()) // 1
}
