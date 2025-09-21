package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi", "pineapple"}

	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println("fruits:", fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Alice", age: 42},
		Person{name: "Bob", age: 32},
		Person{name: "Charlie", age: 22},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("people:", people)

}
