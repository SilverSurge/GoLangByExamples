package main

import "fmt"

func main() {

	// iterate over elements of a slice
	nums := []int{1, 2, 3, 4}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	// iterate over (k,v) or (k) a map
	mappy := map[string]string{"key-a": "value-a", "key-b": "value-b"}

	for k, v := range mappy {
		fmt.Println(k, ":", v)
	}

	for k := range mappy {
		fmt.Println(k)
	}
}
