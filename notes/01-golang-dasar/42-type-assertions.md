# Type assertions

- Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
- Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

```go
func main() {

	result := random()
	resultToString := result.(string)
	fmt.Println(resultToString)

	resultToInt := result.(int) // Panic
	fmt.Println(resultToInt)

}

func random() any {
	return "OK"
}
```

## Type Assertions Menggunakan Switch

- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions

```go
func main() {

	result := random()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)
	}

}

func random() any {
	return true
}
```

Next: [Pointer](./43-pointer.md)
