package main

import "fmt"

type valiadationError struct {
	Message string
}

func (v *valiadationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &valiadationError{"Validation Error"}
	}

	if id != "Bani" {
		return &notFoundError{"Data not Found!!"}
	}

	// OK
	return nil
}

func main() {

	err := saveData("Bani", nil)

	if err != nil {
		// Terjadi error
		if valiadationErr, ok := err.(*valiadationError); ok {
			fmt.Println("ValidationError", valiadationErr.Error())
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("NotFoundError", notFoundErr.Error())

		} else {
			fmt.Println("Unknown Error:", err.Error())
		}

	} else {

		fmt.Println("Success")
	}
}
