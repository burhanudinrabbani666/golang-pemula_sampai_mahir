package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bani := Man{"Bani"}
	bani.Married()

	fmt.Println(bani) // Mr. Bani
}
