package main

import "fmt"

func main() {
	fmt.Println("Channels")
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println("msg:", msg)
}
