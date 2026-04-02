package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String to boolean
	boolean, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	// String to int
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// Int to another base
	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	// Int to String
	stringInt := strconv.Itoa(9999)
	fmt.Println(stringInt)
}
