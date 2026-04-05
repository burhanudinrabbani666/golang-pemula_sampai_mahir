# Channel sebagai parameter

- Dalam kenyataan pembuatan aplikasi, seringnya kita akan mengirim channel ke function lain via parameter
- Sebelumnya kita tahu bahkan di Go-Lang by default, parameter adalah pass by value, artinya value akan diduplikasi lalu dikirim ke function parameter, sehingga jika kita ingin mengirim data asli, kita biasa gunakan pointer (agar pass by reference).
- Berbeda dengan Channel, kita tidak perlu melakukan hal tersebut

```go
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Burhanudin Rabbani"
}

func TestChannelAsParameter(t *testing.T) {

	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}
```

Next: [Channel in out](./9-channel-in-out.md)
