package main

import "fmt"


func sayHello() {
	fmt.Println("Hello!")
}

func main() {
	go sayHello()
	fmt.Println("Main function")
}
