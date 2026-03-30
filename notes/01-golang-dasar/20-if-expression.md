# if expression

- If adalah salah satu kata kunci yang digunakan untuk percabangan
- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
- Hampir di semua bahasa pemrograman mendukung if expression

```go
func main() {
	name := "Bani"

	if name == "Bani" {
		fmt.Println("Hello", name)
	}

}
```

## Else Expression

- Blok if akan dieksekusi ketika kondisi if bernilai true
- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
- Hal ini bisa dilakukan menggunakan else expression

```go
func main() {
	name := "Banis"

	if name == "Bani" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

}
```

## Else If Expression

- Kadang dalam If, kita butuh membuat beberapa kondisi
- Kasus seperti ini, kita bisa menggunakan Else If expression

```go
func main() {
	name := "Agus"

	if name == "Bani" {
		fmt.Println("Hello", name)
	} else if name == "Agus" {
		fmt.Println("Hidup", name)
	} else {
		fmt.Println("Hi, apakah mau coklat?")
	}

}
```

## If dengan Short Statement

- If mendukung short statement sebelum kondisi
- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan terhadap kondisi

```go
func main() {
	name := "Agus salim"

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
```

Next: [Switch expression](./21-switch-expression.md)
