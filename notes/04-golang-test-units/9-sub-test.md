# Sub test

- Go-Lang mendukung fitur pembuatan function unit test di dalam function unit test
- Fitur ini memang sedikit aneh dan jarang sekali dimiliki di unit test di bahasa pemrograman yang lainnya
- Untuk membuat sub test, kita bisa menggunakan function Run()

```go
func TestSubTest( t *testing.T){

	t.Run("Bani", func(t *testing.T) {

		result := HelloWorld("Bani")
		require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	})

	t.Run("Udin", func(t *testing.T) {

		result := HelloWorld("Udin")
		require.Equal(t, "Hi Udin", result, "Result must be Hello Bani")

	})

}
```

## Menjalankan Hanya Sub Test

- Kita sudah tahu jika ingin menjalankan sebuah unit test function, kita bisa gunakan perintah: go test -run TestNamaFunction
- Jika kita ingin menjalankan hanya salah satu sub test, kita bisa gunakan perintah: go test -run TestNamaFunction/NamaSubTest
- Atau untuk semua test semua sub test di semua function, kita bisa gunakan perintah: go test -run /NamaSubTest

Next: [Table test](./10-table-test.md)
