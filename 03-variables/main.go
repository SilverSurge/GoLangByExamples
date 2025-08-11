package main

import "fmt"

func main() {
	// one variable
	var a int = 10
	fmt.Println("a =", a)

	// multiple variables
	var b, c float64 = 20.5, 30.5
	fmt.Println("b =", b, "c =", c)

	// automatic type inference
	var d = true
	fmt.Println("d =", d)

	// shorthand declaration
	e := "Hello, Go!"
	fmt.Println("e =", e)
}
