package main

import "fmt"

func main() {
	// Multiple variables can be declared at once
	var a, b, c int
	fmt.Println(a, b, c)

	// Variables can be initialized at the time of declaration
	var d, e, f int = 1, 2, 3
	fmt.Println(d, e, f)

	// Go can infer the type of initialized variables
	var g, h, i = 4, 5, 6
	fmt.Println(g, h, i)

	// Short hand declaration for variables
	j, k, l := 7, 8, 9
	fmt.Println(j, k, l)

	j, k, l = 10, 11, 12
	fmt.Println(j, k, l)

	nums := []int{1, 2, 3, 4}
	for m, num := range nums {
		fmt.Printf("Index: %d, Value: %d\n", m, num)
	}
}
