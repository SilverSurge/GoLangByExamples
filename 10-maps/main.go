package main

import (
	"fmt"
	"maps"
)

func main() {
	// create a map
	m := make(map[string]int)

	// update values
	m["k1"] = 7
	m["k2"] = 11

	fmt.Println("m:", m)

	// retrieve value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// retrieve value that doesn't exist
	// 0 like value
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// length of map
	fmt.Println("len:", len(m))

	// delete key
	delete(m, "k2")
	fmt.Println("m:", m)

	// clear all elements
	clear(m)
	fmt.Println("map:", m)

	// optional return is_present
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declaration wih values
	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n1:", n1)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n2) {
		fmt.Println("n1 == n2")
	}
}
