package main

import "fmt"

func main() {
	fmt.Println()

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		is_odd := n & 1
		if is_odd != 0 {
			fmt.Println("range", n)
		}
	}
}
