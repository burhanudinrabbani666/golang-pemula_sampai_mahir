package main

import "fmt"

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Bani", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {

	if blacklist(name) {
		fmt.Println("Your blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}

}
