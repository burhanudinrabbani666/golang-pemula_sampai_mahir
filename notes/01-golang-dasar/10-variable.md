# Variable

- Variable adalah tempat untuk menyimpan data
- Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya

```go
func main() {
	var firstName string = "Burhanudin"
	fmt.Println(firstName)

	firstName = "Rabbani" // Ubah Value
	fmt.Println(firstName)

}
```

## Tipe Data Variable

- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib menyebutkan tipe data variable nya

```go
func main() {
	var firstName = "Burhanudin" // set Default ke string
	fmt.Println(firstName)

	firstName = 2 // Error
	fmt.Println(firstName)

}
```

## Kata Kunci Var

- Di Go-Lang, kata kunci var saat membuat variable tidak lah wajib.
- Asalkan saat membuat variable kita langsung menginisialisasi datanya
- Agar tidak perlu menggunakan kata kunci var, kita perlu menggunakan kata kunci := saat menginisialisasikan data pada variable tersebut

```go
func main() {
	firstName := "Burhanudin Rabbani"
	fmt.Println(firstName)

	firstName = "Burhanudin"
	fmt.Println(firstName)
}
```

## Deklarasi Multiple Variable

- Di Go-Lang kita bisa membuat variable secara sekaligus banyak
- Code yang dibuat akan lebih bagus dan mudah dibaca

```go
func main() {
	var (
		firstName  string = "Burhanudin"
		middleName string = "D"
		lastName   string = "Rabbani"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}
```

Next: [constant](./11-constant.md)
