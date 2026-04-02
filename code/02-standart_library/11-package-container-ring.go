package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for index := 0; index < data.Len(); index++ {
		data.Value = "value " + strconv.Itoa(index+1)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
