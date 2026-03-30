# Named return value

- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya mendeklarasikan tipe data return value di function
- Namun kita juga bisa membuat variable secara langsung di tipe data return function nya

```go
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Burhanduin"
	middleName = "D"
	lastName = "Rabbani"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
```

Next: [Variadic function](./29-variadic-function.md)
