package main

import "fmt"

func Ups() any {
	return 1
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)

}
