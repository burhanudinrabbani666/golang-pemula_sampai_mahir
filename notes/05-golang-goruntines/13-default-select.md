# Default select

- Apa yang terjadi jika kita melakukan select terhadap channel yang ternyata tidak ada datanya?
- Maka kita akan menunggu sampai data ada
- Kadang mungkin kita ingin melakukan sesuatu jika misal semua channel tidak ada datanya ketika kita melakukan select channel
- Dalam select, kita bisa menambahkan default, dimana ini akan dieksekusi jika memang di semua channel yang kita select tidak ada datanya

```go
func TestSelectChannelDefault(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {

		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++

		default:
			fmt.Println("Menunggu Data...")
		}

		if counter == 2 {
			break
		}
	}
}
```

Next: [Race condition](./14-race-condition.md)
