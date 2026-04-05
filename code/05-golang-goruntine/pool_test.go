package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool sync.Pool = sync.Pool{
		New: func() any {
			return "New"
		},
	}

	pool.Put("Burhanudin")
	pool.Put("D")
	pool.Put("Rabbani")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)

			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
