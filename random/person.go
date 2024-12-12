package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p :=Person{FirstName: "John", LastName: "Doe"}
	fmt.Println(p.FullName())
}
