package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 // must be 64-bit aligned
	var wg sync.WaitGroup

	numWorkers := 5
	numIncrements := 1000

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numIncrements; j++ {
				atomic.AddInt64(&counter, 1) // safe concurrent increment
				time.Sleep(time.Millisecond) // simulate work
			}
			fmt.Println("Worker", id, "done")
		}(i)
	}

	// No mutexes are needed, and no race conditions occur.

	wg.Wait()

	// Read the counter safely
	final := atomic.LoadInt64(&counter)
	fmt.Println("Final counter value:", final)
}
