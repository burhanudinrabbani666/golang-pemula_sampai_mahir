package main

import "fmt"

func main() {
	names := [...]string{"Bani", "Burhan", "Udin", "Ani", "D", "Bandot"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[1:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)

	// Append ------------------------

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"  // Ini mengubah data days
	daysSlice1[1] = "Minggu Baru" // Ini mengubah data days

	daysSlice2 := append(daysSlice1, "Libur Baru")
	daysSlice2[0] = "Sabtu Lama"

	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)

	// Make ------------------

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bani"
	newSlice[1] = "Bani"
	// newSlice[2] = "Bani" // Error, harus nya dengan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Bani")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Heri"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	// Copy ------------------

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice, toSlice)

	// Array vs Slice -------

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray, iniSlice)

}
