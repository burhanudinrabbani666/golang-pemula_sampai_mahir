package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Bani",
		"address": "cirebon",
	}

	fmt.Println(person["name"])    // Bani
	fmt.Println(person["address"]) // cirebon
	fmt.Println(person)            // map[address:cirebon name:Bani]

	// ----------------------

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Bani"
	book["ups"] = "Salah"

	fmt.Println(book)
	delete(book, "ups")
	fmt.Println(book)

}
