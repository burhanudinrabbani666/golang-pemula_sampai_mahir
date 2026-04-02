package helper

import "testing"

func TestHelloWorld(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		panic("Result is not Hello Bani")
	}
}

