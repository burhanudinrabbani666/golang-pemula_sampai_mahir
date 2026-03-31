package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animel Animal) GetName() string {
	return animel.Name
}

func main() {
	person := Person{Name: "Bani"}
	sayHello(person)

	animal := Animal{Name: "Cat"}
	sayHello(animal)
}
