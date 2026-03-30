# Struct

- Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
- Struct biasanya representasi data dalam program aplikasi yang kita buat
- Data di struct disimpan dalam field
- Sederhananya struct adalah kumpulan dari field

## Membuat Data Struct

- Struct adalah template data atau prototype data
- Struct tidak bisa langsung digunakan
- Namun kita bisa membuat data/object dari struct yang telah kita buat

```go

type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// 1.
	var bani Customer

	bani.Name = "Burhanudin"
	bani.Address = "Cirebon"
	bani.Age = 23

	// 2.
	nico := Customer{
		Name:    "Nico",
		Address: "Cirebon",
		Age:     24,
	}

	// 3.
	anang := Customer{"Anang", "Bekasi", 25}

}
```

Next: [Struct method](./38-struct-method.md)
