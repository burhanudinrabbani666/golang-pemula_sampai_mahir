# Operator new

- Sebelumnya untuk membuat pointer dengan menggunakan operator &
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal

```go
func main() {

	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
```

Next: [Pointer di function](./46-pointer-di-function.md)
