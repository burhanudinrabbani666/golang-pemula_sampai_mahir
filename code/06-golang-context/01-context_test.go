package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContecxt(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f")) // F
	fmt.Println(contextF.Value("c")) // C --> Ambil Ke Parent
	fmt.Println(contextF.Value("b")) // <nil> Beda Parent
	fmt.Println(contextA.Value("b")) // <nil> Tidak bisa mengambil child

	fmt.Println(contextE.Value("d")) // <nil> tidak bisa ke siblings
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return

			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulasi Slow
			}
		}

	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			break
		}
	}

	cancel() // Mengirim sinyal cancel ke context

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

}

func TestContextWithTimeout(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

}

func TestContextWithDeadline(t *testing.T) {
	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2
	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3

	for n := range destination {
		fmt.Println("Counter:", n)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

}
