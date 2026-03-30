# Struct method

- Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
- Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
- Method adalah function

```go
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My name is", customer.Name)
}

func main() {
	bani := Customer{"Bani", "Cirebon", 23}

	bani.sayHello()
}
```

Next: [Interface](./39-interface.md)
