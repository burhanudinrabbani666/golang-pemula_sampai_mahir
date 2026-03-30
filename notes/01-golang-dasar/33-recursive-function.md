# Recursive function

- Recursive function adalah function yang memanggil function dirinya sendiri
- Kadang dalam pekerjaan, kita sering menemui kasus dimana menggunakan recursive function lebih mudah dibandingkan tidak menggunakan recursive function
- Contoh kasus yang lebih mudah diselesaikan menggunakan recursive adalah Factorial

```go
func factorailLoop(value int) int {
	result := 1

	for index := value; index > 0; index-- {
		result *= index
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	result := 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
	fmt.Println(result)

	fmt.Println(factorailLoop(10))
	fmt.Println(factorialRecursive(10))
}
```

Next: [Closure](./34-closure.md)
