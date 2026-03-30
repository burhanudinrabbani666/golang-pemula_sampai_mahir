# For expression

- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
- Salah satu fitur perulangan adalah for loops

```go
func main() {
	counter := 1

	for counter <= 10 { // Perulangan dulu sampai selesai
		fmt.Println("Perulangan ke: ", counter)
		counter++
	}

	fmt.Println("Selesai")
}
```

## For dengan Statement

- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan di for
- Init statement, yaitu statement sebelum for di eksekusi
- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan

```go
func main() {

    //   init                          post
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke: ", counter)
	}

	fmt.Println("Selesai")
}
```

## For Range

- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
- Data collection contohnya Array, Slice dan Map

```go
	names := []string{"Bani", "Agus", "Ryan", "Heri"}
	for _, name := range names {
		fmt.Println(name)
	}
```

Next: [Break and Continue](./23-break-continue.md)
