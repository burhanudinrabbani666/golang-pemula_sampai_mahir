package helper

import (
	"fmt"
	"testing"
)

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
