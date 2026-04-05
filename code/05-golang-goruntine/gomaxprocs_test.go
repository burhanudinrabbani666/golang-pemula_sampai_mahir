package golanggoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomxprocs(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Go(func() {
			time.Sleep(3 * time.Second)
		})
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoRoutine)

	group.Wait()
}

func TestChangeThreadNum(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Go(func() {
			time.Sleep(3 * time.Second)
		})
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	runtime.GOMAXPROCS(20) // Jarang karna golang sudah optimal
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoRoutine)

	group.Wait()
}
