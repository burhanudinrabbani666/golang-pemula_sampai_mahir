package main

import "fmt"

func endApp() {
	fmt.Println("End Application")

	message := recover()
	fmt.Println("Terjadi Error", message)

}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups ERROR!!!")
	}

}

func main() {
	runApp(true)

	fmt.Println("Burhanudin Rabbani")
}
