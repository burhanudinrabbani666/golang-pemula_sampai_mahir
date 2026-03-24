# Tipe data string

- String ada tipe data kumpulan karakter
- Jumlah karakter di dalam String bisa nol sampai tidak terhingga
- Tipe data String di Go-Lang direpresentasikan dengan kata kunci string
- Nilai data String di Go-Lang selalu diawali dengan karakter “ (petik dua) dan diakhiri dengan karakter “ (petik dua)

```go
func main() {
	fmt.Println("Burhanudin Rabbani")
}

```

| Function         | Keterangan                                     |
| ---------------- | ---------------------------------------------- |
| len(“string”)    | Menghitung jumlah karakter di String           |
| “string”[number] | Mengambil karakter pada posisi yang ditentukan |

```go
func main() {
	fmt.Println(len("Bani")) // 4
	fmt.Println("Bani"[0])   //  66
}
```

Kenapa 66? Kita dapat Byte bukan charcter

Next: [Variable](./10-variable.md)
