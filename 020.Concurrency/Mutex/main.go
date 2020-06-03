package main

import (
	"fmt"
	"runtime"
	"sync"
)

var WG sync.WaitGroup

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100

	WG.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++

			counter = v
			mu.Unlock()
			WG.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	WG.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
