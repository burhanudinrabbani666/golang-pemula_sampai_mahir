package main

import (
	"fmt"
	"golang-pemula_sampai_mahir/database"
	_ "golang-pemula_sampai_mahir/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
