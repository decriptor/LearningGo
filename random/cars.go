package main

import "fmt"

type Engine struct {
	Horsepower int
}

type Car struct {
	Engine // Embedding Engine in Car
	Brand string
}

func main() {
	car := Car{
		Engine: Engine{Horsepower: 150},
		Brand: "Toyota",
	}
	fmt.Println(car.Brand, "with", car.Horsepower, "HP")
}

