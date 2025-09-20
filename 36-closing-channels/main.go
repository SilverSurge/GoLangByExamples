package main

import "fmt"

func main() {

	// make two channels
	jobs := make(chan int, 5)
	done := make(chan bool)

	// anonymous go routine
	go func() {
		for {
			// this special 2-value form of receive,
			//  the more value will be false
			// if jobs has been closed and
			// all values in the channel have already been received
			j, more := <-jobs
			if more {
				fmt.Println("recieved job", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 4; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("recieved more jobs:", ok)
}
