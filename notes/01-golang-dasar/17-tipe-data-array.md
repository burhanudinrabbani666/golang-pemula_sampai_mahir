# Tipe data array

- Array adalah tipe data yang berisikan kumpulan data dengan tipe yang sama
- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
- Daya tampung Array tidak bisa bertambah setelah Array dibuat

Index di Array

| Index | Data      |
| ----- | --------- |
| 0     | Eko       |
| 1     | Kurniawan |
| 2     | Khannedy  |

```go
func main() {
	var names [3]string

	names[0] = "Burhanudin"
	names[1] = "D"
	names[2] = "Rabbani"

	fmt.Println(names)

}
```

Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable

```go
func main() {

	var values = [3]int{12, 23, 34}

	fmt.Println(values)
}
```

| Operasi              | Keterangan                      |
| -------------------- | ------------------------------- |
| len(array)           | Untuk mendapatkan panjang Array |
| array[index]         | Mendapat data di posisi index   |
| array[index] = value | Mengubah data di posisi index   |

Next: [Tipe data slice](./18-tipe-data-slice.md)
