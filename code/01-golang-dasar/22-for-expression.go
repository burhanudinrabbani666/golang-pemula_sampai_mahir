package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Bani", "Agus", "Ryan", "Heri"}
	for _, name := range names {
		fmt.Println(name)
	}

}
