package main

import (
	"fmt"
)

func main() {
	type NoKTP string // Type declaration

	var ktpBani NoKTP = "12345678"

	fmt.Println(ktpBani)
	fmt.Println(NoKTP("87654321"))

}
