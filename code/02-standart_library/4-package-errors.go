package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {

	if id == "" {
		return ValidationError
	}

	if id != "Bani" {
		return NotFoundError
	}

	return nil
}

func main() {

	err := GetById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error:", err)
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error:", err)
		} else {
			fmt.Println("Unknown Error", err)
		}

	}

}
