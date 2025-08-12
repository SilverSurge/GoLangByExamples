package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("s:", s)

	const a = 10000

	const b = 3e20 / a

	fmt.Println("a:", a)

	fmt.Println("sin(a)", math.Sin(a))

	fmt.Println("b:", b)

	fmt.Println("int64(b):", int64(b))
}
