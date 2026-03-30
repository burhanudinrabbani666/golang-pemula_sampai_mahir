package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increnment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
