# Closure

- Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
- Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi

```go
func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increnment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
```

Next: [Defer Panic Recover](./35-defer-panic-recover.md)
