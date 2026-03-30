# Switch expression

- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
- Switch expression sangat sederhana dibandingkan if
- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu variable

```go
func main() {
	name := "Bani"

	switch name {
	case "Bani":
		fmt.Println("Hello Bani")

	case "Agus":
		fmt.Println("Hello Agus")

	case "Heri":
		fmt.Println("Hello Heri")

	default:
		fmt.Println("Hi, Oasis juga a?")
	}
}
```

## Switch dengan Short Statement

Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek kondisinya

```go
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu panjang")

	case false:
		fmt.Println("Nama sudah benar")

	}
```

## Switch Tanpa Kondisi

- Kondisi di switch expression tidak wajib
- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut di setiap case nya

```go
	length := len(name)

	switch{
	case length > 10:
		fmt.Println("nama terlau panjang")
	case length >5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("Nam sudah benar")
	}
```

Next: [For expression](./22-for-expression.md)
