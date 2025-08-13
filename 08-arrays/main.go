package main

import "fmt"

func main() {

	// fixed size array, default 0
	var a [5]int
	fmt.Println("a:", a)

	// set an element
	a[0] = 1
	fmt.Println("a[0]:", a)

	// array with elements
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// set some elements and rest to zero
	c := [10]int{100, 3: 400, 500, 8: 900}
	fmt.Println("c:", c)

	// auto decide length
	d := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println("d:", d)

	// two dimensional array
	var two_d [2][3]int

	for i := 0; i < 2; i++ {
		for j := range 3 {
			two_d[i][j] = i + j
		}
	}

	fmt.Println("two_d:", two_d)
}
