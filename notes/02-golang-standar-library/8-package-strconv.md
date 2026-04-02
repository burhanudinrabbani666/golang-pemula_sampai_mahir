# Package strconv

- Sebelumnya kita sudah belajar cara konversi tipe data, misal dari int32 ke int34
- Bagaimana jika kita butuh melakukan konversi yang tipe datanya berbeda? Misal dari int ke string, atau sebaliknya
- Hal tersebut bisa kita lakukan dengan bantuan package strconv (string conversion)
- https://golang.org/pkg/strconv/

## Beberapa Function di Package strconv

| Function                       | Kegunaan                   |
| ------------------------------ | -------------------------- |
| strconv.parseBool(string)      | Mengubah string ke bool    |
| strconv.parseFloat(string)     | Mengubah string ke float   |
| strconv.parseInt(string)       | Mengubah string ke int64   |
| strconv.FormatBool(bool)       | Mengubah bool ke string    |
| strconv.FormatFloat(float, … ) | Mengubah float64 ke string |
| strconv.FormatInt(int, … )     | Mengubah int64 ke string   |

```go
func main() {

	// String to boolean
	boolean, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(boolean)
	}

	// String to int
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// Int to another base
	binary := strconv.FormatInt(999, 2)
	fmt.Println(binary)

	// Int to String
	stringInt := strconv.Itoa(9999)
	fmt.Println(stringInt)
}
```

Next: [Package math](./9-package-math.md)
