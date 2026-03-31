# Pointer di method

- Walaupun method akan menempel di struct, tapi sebenarnya data struct yang diakses di dalam method adalah pass by value
- Sangat direkomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method

```go
type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	bani := Man{"Bani"}
	bani.Married()

	fmt.Println(bani) // Mr. Bani
}
```

Next: [Package import](./48-package-import.md)
