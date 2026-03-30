# Anonymous function

- Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
- Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus membuat function terlebih dahulu
- Hal tersebut dinamakan anonymous function, atau function tanpa nama

```go
func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Bani", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {

	if blacklist(name) {
		fmt.Println("Your blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}

}
```

Next: [Recursive function](./33-recursive-function.md)
