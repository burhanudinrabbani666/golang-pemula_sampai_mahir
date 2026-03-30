# Break continue

- Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
- Break digunakan untuk menghentikan seluruh perulangan
- Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

```go
func main() {
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke:", i)
	}
}
```

```go
func main() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke:", i)
	}
}
```

Next: [function](./24-function.md)
