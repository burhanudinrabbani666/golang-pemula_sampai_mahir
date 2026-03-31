# Pointer di function

- Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di copy lalu dikirim ke function tersebut
- Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah berubah.
- Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
- Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
- Untuk melakukan ini, kita juga bisa menggunakan pointer di function
- Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator \* di parameternya

```go
func changeCountryToIndonesia(address *Address) { // Pakai * di parameternya
	address.Country = "Indonesia"
}

func main() {

	address := Address{}
	changeCountryToIndonesia(&address) // Pakai & ketika input parameternya

	fmt.Println(address)
}
```

Next: [](./47-pointer-di-method.md)
