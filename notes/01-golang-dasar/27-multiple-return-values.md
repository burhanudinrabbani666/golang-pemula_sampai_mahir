# Multiple return values

- Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
- Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe data return value nya di function

```go
func getFullName() (string, string) {
	return "Burhanudin", "Rabbani"
}

func main() {
	firstName, lastName := getFullName() // Bisa namakan apa saja

	fmt.Println(firstName, lastName)
}

```

## Menghiraukan Return Value

- Multiple return value wajib ditangkap semua value nya
- Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda \_ (garis bawah)

```go
func main() {
	firstName, _ := getFullName() // Bisa namakan apa saja

	fmt.Println(firstName, _)
}
```

Next: [Named return value](./28-named-return-value.md)
