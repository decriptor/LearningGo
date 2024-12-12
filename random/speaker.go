package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(s.Speak())
}
