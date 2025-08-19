package main

import "fmt"

func intSeq() func() int {
	// this gets attached to the scope of the function
	i := 0
	return func() int {
		// acts like a generator
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
