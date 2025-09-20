package main

import (
	"fmt"
	"time"
)

func main() {

	// create requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// `limiter` channel will recieve a value every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// allows a burst of 3
	burstyLimiter := make(chan time.Time, 3)

	for range 3 {
		burstyLimiter <- time.Now()
	}

	// refill the `burstLimiter`
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	fmt.Println()
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
