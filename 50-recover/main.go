package main

import "fmt"

func worker1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("worker1 recovery!!")
		}
	}()
	panic("boom!!")
}

func worker2() {
	panic("blasted!!")
}

func main() {
	go worker1()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recovery!!")
		}
	}()

	worker2()
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func mayPanic() {
// 	panic("a problem")
// }

// func main() {
// 	go func() {
// 		for {
// 			time.Sleep(time.Second)
// 			fmt.Println("slept for a second")
// 		}
// 	}()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("recovered error\n", r)
// 			go func() {
// 				for {
// 					time.Sleep(time.Second)
// 					fmt.Println("slept for a second in recover block")
// 				}
// 			}()

// 			time.Sleep(time.Second * 4)
// 		}
// 	}()

// 	time.Sleep(time.Second * 4)

// 	mayPanic()

// 	fmt.Println("after panic")
// }
