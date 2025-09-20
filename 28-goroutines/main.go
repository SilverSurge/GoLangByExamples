package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func main() {
	fmt.Println("Goroutines")

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// better approach would be to use waitgroup
	time.Sleep(time.Second)
	fmt.Println("done")
}
