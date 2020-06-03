package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wG sync.WaitGroup

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	wG.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++

			counter = v
			wG.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wG.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
