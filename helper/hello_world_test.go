package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Bani")
	require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	fmt.Println("TestHelloWorld with Assert Done")
}


func TestHelloWorldAssertion(t *testing.T){
	result := HelloWorld("Bani")
	assert.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	fmt.Println("TestHelloWorld with Assert Done")
}


func TestHelloWorld(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Fail()
	}

	fmt.Println("Ini test Hello World")
}

func TestHelloWorldBani(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Error("Result must be Hello Bani") // Ini akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Dieksekusi walaupun Error")

}


func TestHelloWorldUdin(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Fatal("Result must be Hello Bani") // Ini tidak akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Tidak dieksekusi ketika Error")
}