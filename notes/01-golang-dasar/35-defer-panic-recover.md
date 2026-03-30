# Defer Panic Recover

## Defer

- Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
- Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

```go
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()

	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
```

## Panic

- Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

```go
func endApp() {
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups ERROR!!!")
	}
}

func main() {
	runApp(true)
}

```

## Recover

- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

**Cara salah ❎**

```go
func endApp() {
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups ERROR!!!")
	}

	message := recover()
	fmt.Println("Terjadi Error", message)
}
```

**Cra Benar ✅**

```go
func endApp() {
	fmt.Println("End Application")

	message := recover()
	fmt.Println("Terjadi Error", message)

}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups ERROR!!!")
	}

}

func main() {
	runApp(true)

	fmt.Println("Burhanudin Rabbani")
}
```

Next: []()
