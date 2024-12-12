package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Send
	}()

	val := <- ch // Recieve
	fmt.Println(val)
}
