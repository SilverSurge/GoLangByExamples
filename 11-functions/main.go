package main

import "fmt"

// return type int
func plus(a int, b int) int {

	return a + b
}

// a, b, and c get type int
func plusPlus(a, b, c int) int {
	return a + b + c
}

// a, b - int
// w1, w2 - string
func plusPrint(a, b int, w1, w2 string) int {
	fmt.Printf("%s + %s\n", w1, w2)
	return a + b
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = plusPrint(1, 10, "one", "ten")
	fmt.Println("1+10 =", res)
}
