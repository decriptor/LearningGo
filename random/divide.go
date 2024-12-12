package main

import "fmt"

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}


func main() {
	result, error := divide(4, 5)
	fmt.Println(result)
	fmt.Println(error)
}
