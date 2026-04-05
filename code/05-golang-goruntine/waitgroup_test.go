package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronus(group *sync.WaitGroup) {
	defer group.Done() // Kalau tidak ada ini bakal error deadlock

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go RunAsynchronus(group)
	}

	group.Wait()
	fmt.Println("Selesai")
}

// Once
