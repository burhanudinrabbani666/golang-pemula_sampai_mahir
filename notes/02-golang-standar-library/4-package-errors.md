# Package errors

- Sebelumnya kita sudah membahas tentang interface error yang merupakan representasi dari error di Go-Lang, dan membuat error menggunakan function errors.New()
- Sebenarnya masih banyak yang bisa kita lakukan menggunakan package errors, contohnya ketika kita ingin membuat beberapa value error yang berbeda
- https://pkg.go.dev/errors

## Mengecek Jenis Error

- Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
- Kita bisa menggunakan errors.Is() untuk mengecek jenis type error nya

```go
var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {

	if id == "" {
		return ValidationError
	}

	if id != "Bani" {
		return NotFoundError
	}

	return nil
}

func main() {

	err := GetById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error:", err)
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error:", err)
		} else {
			fmt.Println("Unknown Error", err)
		}

	}

}

```

Next: [Package os](./5-package-os.md)
