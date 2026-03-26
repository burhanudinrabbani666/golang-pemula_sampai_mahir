# Konversi tipe data

- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
- Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

```go
func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32) // -32768? Number overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
}
```

```go
func main() {
	var name = "Burhanudin Rabbani"
	var e = name[0] // uint8
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)       // 66 => uint8
	fmt.Println(eString) // B

}

```

Next: [Type declaration](./13-type-declaration.md)
