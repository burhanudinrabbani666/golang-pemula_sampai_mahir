# Tipe data slice

- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- **Slide dan Array selalu terkoneksi**, dimana Slice adalah data yang mengakses sebagian atau seluruh data di Array

## Detail Tipe Data Slice

- Tipe Data Slice memiliki 3 data, yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array para slice
- Length adalah panjang dari slice, dan
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

## Membuat Slice Dari Array

| Membuat Slice   | Keterangan                                                             |
| --------------- | ---------------------------------------------------------------------- |
| array[low:high] | Membuat slice dari array dimulai index low sampai index sebelum high   |
| array[low:]     | Membuat slide dari array dimulai index low sampai index akhir di array |
| array[:high]    | Membuat slice dari array dimulai index 0 sampai index sebelum high     |
| array[:]        | Membuat slice dari array dimulai index 0 sampai index akhir di array   |

## Function Slice

| Operasi                            | Keterangan                                                                                                                 |
| ---------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| len(slice)                         | Untuk mendapatkan panjang                                                                                                  |
| cap(slice)                         | Untuk mendapat kapasitas                                                                                                   |
| append(slice, data)                | Membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru |
| make([]TypeData, length, capacity) | Membuat slice baru                                                                                                         |
| copy(destination, source)          | Menyalin slice dari source ke destination                                                                                  |

```go
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
}
```

> Saat membuat Array, kita harus berhati-hati, jika salah, maka yang kita buat bukanlah Array, melainkan Slice

```go
	// Array vs Slice -------

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray, iniSlice)

```

Next: [Tipe data map](./19-tipe-data-map.md)
