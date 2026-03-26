# Type declaration

- Type Declarations adalah kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
- Type Declarations biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, dengan tujuan agar lebih mudah dimengerti

```go
func main() {
	type NoKTP string // Type declaration

	var ktpBani NoKTP = "12345678"

	fmt.Println(ktpBani)
	fmt.Println(NoKTP("87654321"))
}
```

Next: [Operasi matematika](./14-operasi-matematika.md)
