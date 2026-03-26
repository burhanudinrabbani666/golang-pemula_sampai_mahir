package main

import "fmt"

func main() {
	var name = "Burhanudin Rabbani"
	var e = name[0] // uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)       // 66 => uint8
	fmt.Println(eString) // B

}
