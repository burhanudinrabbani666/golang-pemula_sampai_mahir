# Package strings

- Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data String
- Ada banyak sekali function yang bisa kita gunakan
- https://golang.org/pkg/strings/

## Beberapa Function di Package strings

| Function                             | Kegunaan                                         |
| ------------------------------------ | ------------------------------------------------ |
| strings.Trim(string, cutset)         | Memotong cutset di awal dan akhir string         |
| strings.ToLower(string)              | Membuat semua karakter string menjadi lower case |
| strings.ToUpper(string)              | Membuat semua karakter string menjadi upper case |
| strings.Split(string, separator)     | Memotong string berdasarkan separator            |
| strings.Contains(string, search)     | Mengecek apakah string mengandung string lain    |
| strings.ReplaceAll(string, from, to) | Mengubah semua string dari from ke to            |

```go
func main() {
	fmt.Println(strings.Contains("Burhanudin Rabbani", "bani"))    // true
	fmt.Println(strings.Split("Burhanudin Rabbani Aurelius", " ")) // []string
	fmt.Println(strings.ToLower("Burhanudin Rabbani Aurelius"))
	fmt.Println(strings.ToUpper("Burhanudin Rabbani Aurelius"))
	fmt.Println(strings.Trim("            Burhanudin Rabbani Aurelius             ", " "))
	fmt.Println(strings.ReplaceAll("Burhanudin Rabbani Aurelius bani", "bani", "respect"))

}
```

Next: [Package strconv](./8-package-strconv.md)
