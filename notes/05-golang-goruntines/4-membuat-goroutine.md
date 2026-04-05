# Membuat goroutine

- Untuk membuat goroutine di Golang sangatlah sederhana
- Kita hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai

```go
func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld() // Menambahkan go untuk async
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}
```

Next: [Goroutine ringan](./5-goroutine-ringan.md)
