package main

import "fmt"

func main() {
	nilaiAkhir := 90
	absensi := 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) // false
}
