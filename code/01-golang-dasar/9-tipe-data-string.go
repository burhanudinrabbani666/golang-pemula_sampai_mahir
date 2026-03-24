package main

import "fmt"

func main() {
	nama := "Burhanudin Rabbani"
	fmt.Println(nama)

	teks := "Bani"
	fmt.Println(len(teks)) // 4 (jumlah byte)

	// Akses per byte (bukan per karakter/rune).
	bytePertama := teks[0]
	fmt.Println(bytePertama)        // 66 (nilai byte huruf 'B')
	fmt.Printf("%c\n", bytePertama) // B
}
