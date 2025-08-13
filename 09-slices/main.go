package main

import (
	"fmt"
	"slices"
)

func main() {
	// create an empty slice
	var s []string
	fmt.Println("s:", s)
	fmt.Println("s==nil:", s == nil)
	fmt.Println("len(s):", len(s))

	// make a slice of length 3
	s = make([]string, 3)
	fmt.Println("\ns:", s)
	fmt.Println("s==nil:", s == nil)
	fmt.Println("len(s):", len(s))

	// modify slices like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	// append to slices
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("\ns:", s)
	fmt.Println("s==nil:", s == nil)
	fmt.Println("len(s):", len(s))

	// copy slices
	t := make([]string, len(s))
	copy(t, s)

	// slicing slices
	// similar to python slicing
	fmt.Println("\nt[1:3]", t[1:3])

	// check if equal
	if slices.Equal(s, t) {
		fmt.Println("s == t")
	}

	// two dimensional slices
	slice_2d := make([][]int, 3)
	for i := range 3 {
		sz := i + 1
		slice_2d[i] = make([]int, sz)
		for j := range sz {
			slice_2d[i][j] = i + j
		}
	}
	fmt.Println("\nslice_2d:", slice_2d)
}
