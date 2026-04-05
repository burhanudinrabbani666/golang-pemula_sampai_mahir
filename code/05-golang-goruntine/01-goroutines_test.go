package golanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld() // Menambahkan go untuk async
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
