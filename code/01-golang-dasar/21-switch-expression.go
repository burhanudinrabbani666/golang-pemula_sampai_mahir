package main

import "fmt"

func main() {
	name := "Bani"

	switch name {
	case "Bani":
		fmt.Println("Hello Bani")

	case "Agus":
		fmt.Println("Hello Agus")

	case "Heri":
		fmt.Println("Hello Heri")

	default:
		fmt.Println("Hi, Oasis juga a?")
	}

	// -------------------------------

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu panjang")

	case false:
		fmt.Println("Nama sudah benar")

	}

	// ------------------

}
