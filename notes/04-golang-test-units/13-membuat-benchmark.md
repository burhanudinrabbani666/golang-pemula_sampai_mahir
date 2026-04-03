# Membuat benchmark

## Benchmark Function

- Mirip seperti unit test, untuk benchmark pun, di Go-Lang sudah ditentukan nama function nya, harus diawali dengan kata Benchmark, misal BenchmarkHelloWorld, BenchmarkXxx
- Selain itu, harus memiliki parameter (b \*testing.B)
- Dan tidak boleh mengembalikan return value
- Untuk nama file benchmark, sama seperti unit test, diakhiri dengan \_test, misal hello_world_test.go

## Menjalankan Benchmark

- Untuk menjalankan seluruh benchmark di module, kita bisa menggunakan perintah sama seperti test, namun ditambahkan parameter bench :

```bash
go test -v -bench=.
```

- Jika kita hanya ingin menjalankan benchmark tanpa unit test, kita bisa gunakan perintah :

```bash
go test -v -run=NotMathUnitTest -bench=.
```

- Kode diatas selain menjalankan benchmark, akan menjalankan unit test juga, jika kita hanya ingin menjalankan benchmark tertentu, kita bisa gunakan perintah :

```bash
go test -v -run=NotMathUnitTest -bench=BenchmarkTest
```

- Jika kita menjalankan benchmark di root module dan ingin semua module dijalankan, kita bisa gunakan perintah :

```bash
go test -v -bench=. ./...
```

```go
func BenchmarkHelloWorld(b *testing.B) {
	for index := 0; index < b.N; index++ {
		HelloWorld("Bani")
	}
}

func BenchmarkHelloWorldBurhanudin(b *testing.B) {
	for index := 0; index < b.N; index++ {
		HelloWorld("Burhanudin")
	}
}
```

Next: [Sub benchmark](./14-sub-benchmark.md)
