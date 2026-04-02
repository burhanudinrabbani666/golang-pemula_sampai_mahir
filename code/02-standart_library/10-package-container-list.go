package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Burhanudin")
	data.PushBack("Rabbani")
	data.PushBack("Aurelius")

	// --------------------------------------------------------- //
	//                       Manual                              //
	// --------------------------------------------------------- //

	var head *list.Element = data.Front() // Burhanudin
	fmt.Println(head.Value)

	next := head.Next() // Rabbani
	fmt.Println(next.Value)

	next = next.Next()
	fmt.Println(next.Value)

	// --------------------------------------------------------- //
	//                      For Loop                             //
	// --------------------------------------------------------- //

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
