# Package import

- Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di Go-Lang
- Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
- Package sendiri sebenarnya hanya direktori folder di sistem operasi kita

```go
// helper/helper.go
package helper

func SayHello(name string) string {
	return "Hello " + name

}
```

```go
import (
	"fmt"
	"golang-pemula_sampai_mahir/helper"
)

func main() {

	result := helper.SayHello("Bani")
	fmt.Println(result)
}
```

Next: [Access modifier](./49-access-modifier.md)
