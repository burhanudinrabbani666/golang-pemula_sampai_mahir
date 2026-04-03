# Sub benchmark

Sama seperti testing.T, di testing.B juga kita bisa membuat sub benchmark menggunakan function Run()

```go
func BenchmarkSub(b *testing.B) {
	b.Run("Bani", func(b *testing.B) {
		for index := 0; index < b.N; index++ {
			HelloWorld("Bani")
		}
	})

	b.Run("Burhanudin", func(b *testing.B) {
		for index := 0; index < b.N; index++ {
			HelloWorld("Burhanudin")
		}
	})

}
```

Menjalankan Hanya Sub Benchmark

- Saat kita menjalankan benchmark function, maka semua sub benchmark akan berjalan
- Namun jika kita ingin menjalankan salah satu sub benchmark saja, kita bisa gunakan perintah :

```bash
go test -v -bench=BenchmarkNama/NamaSub
```

Next: [Table benchmark](./15-table-benchmark.md)
