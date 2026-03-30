package main

import "fmt"

func getFullName() (string, string) {
	return "Burhanudin", "Rabbani"
}

func main() {
	firstName, _ := getFullName() // Bisa namakan apa saja

	fmt.Println(firstName, _)
}
