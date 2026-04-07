# Embed file ke byte

- Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
- Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain

```go
//go:embed images.png
var images []byte

func TestByte(t *testing.T) {

	err := os.WriteFile("images_new.png", images, fs.ModePerm)

	if err != nil {
		panic(err)
	}

}
```

Next: [Embed multiple file](./04-embed_multiple_file.md)
