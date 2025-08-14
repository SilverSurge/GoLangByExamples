package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	// use both values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use _ to ignore values
	_, c := vals()
	fmt.Println(c)
}
