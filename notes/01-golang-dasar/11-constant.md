# Constant

- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya

> Tidak masalah membuat constant lalu tidak dipakai, berbeda dengan variable (var)

```go
func main() {
	const firstName string = "Burhanudin"
	const lastName = "Rabbani"

	fmt.Println(firstName, lastName)
}
```

atau

```go
func main() {
	const (
		firstName string = "Burhanudin"
		lastName         = "Rabbani"
	)

	fmt.Println(firstName, lastName)
}
```

Next: [Konversi tipe data](./12-konversi-tipe-data.md)
