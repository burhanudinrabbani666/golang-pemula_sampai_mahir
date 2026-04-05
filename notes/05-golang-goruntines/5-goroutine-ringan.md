# Goroutine ringan

- Seperti yang sebelumnya dijelaskan, bahwa goroutine itu sangat ringan
- Kita bisa membuat ribuan, bahkan sampai jutaan goroutine tanpa takut boros memory
- Tidak seperti thread yang ukurannya berat, goroutine sangatlah ringan

```go
func DisplayNumber(number int) {
	fmt.Println("Display: ", number)
}
func TestManyGorutines(t *testing.T) {

	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(10 * time.Second)
}
```

Next: [Pengenalan channel](./6-pengenalan-channel.md)
