package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // Pointer ke address 1

	address2.City = "Bandung" // mengubah address 1

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // Mengacu ke yang Baru

	fmt.Println(address1) // address satu berubah
	fmt.Println(address2)

}
