# Error

error Interface

- Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface nya adalah error

## Membuat Error

- Untuk membuat error, kita tidak perlu manual.
- Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di package errors (Package akan kita bahas secara detail di materi tersendiri)

```go
import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan Nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {

	hasil, error := Pembagian(100, 0)

	if error == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error:", error.Error())
	}
}
```

Next: [Custom error](./52-custom-error.md)
