# Embed file ke string

- Embed file bisa kita lakukan ke variable dengan tipe data string
- Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut

```go
package golangembed

import (
	_ "embed"
	"fmt"
	"os"
	"testing"
)

// Harus diluar function
//
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	// Traditional read File
	result, err := os.ReadFile("version.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))

	// Hasil dari embed
	fmt.Println(version)
}
```

Next: [Embed file ke byte](./03-embed_file_ke_byte.md)
