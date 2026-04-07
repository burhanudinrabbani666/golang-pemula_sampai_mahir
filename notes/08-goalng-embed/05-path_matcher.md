# Path matcher

- Selain manual satu per satu, kita bisa mengguakan patch matcher untuk membaca multiple file yang kita inginkan
- Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
- Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match

```go
//go:embed files/*.txt
var path embed.FS

func TestPathEmbed(t *testing.T) {
	dirEntrys, _ := path.ReadDir("files")

	for _, entry := range dirEntrys {
		if !entry.IsDir() {

			fmt.Println(entry.Name())

			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))

			fmt.Println("---------")
		}
	}

}
```

Next: [Hasil embed](./06-hasil_embed.md)
