# Hasil embed

- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golang nya
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya, dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi

```go
//go:embed version.txt
var version string

//go:embed images.png
var images []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("images_new.png", images, fs.ModePerm)
	if err != nil {
		panic(err)
	}

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
