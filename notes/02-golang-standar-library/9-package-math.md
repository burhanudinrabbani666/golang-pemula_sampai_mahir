# Package math

- Package math merupakan package yang berisikan constant dan fungsi matematika
- https://golang.org/pkg/math/

## Beberapa Function di Package math

| Function                   | Kegunaan                                                                 |
| -------------------------- | ------------------------------------------------------------------------ |
| math.Round(float64)        | Membulatkan float64 keatas atau kebawah, sesuai dengan yang paling dekat |
| math.Floor(float64)        | Membulatkan float64 kebawah                                              |
| math.Ceil(float64)         | Membulatkan float64 keatas                                               |
| math.Max(float64, float64) | Mengembalikan nilai float64 paling besar                                 |
| math.Min(float64, float64) | Mengembalikan nilai float64 paling kecil                                 |

```go
func main() {
	fmt.Println(math.Ceil(1.40))  // 2
	fmt.Println(math.Floor(1.40)) // 1
	fmt.Println(math.Round(1.40)) // 1
	fmt.Println(math.Max(10, 11)) // 11
	fmt.Println(math.Min(10, 11)) // 10
}
```

Next: [Package container list](./10-package-container-list.md)
