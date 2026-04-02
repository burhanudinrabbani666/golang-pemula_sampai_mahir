# Package reflect

- Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode kita pada saat aplikasi sedang berjalan
- Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
- Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video, Anda bisa eksplorasi package reflec ini secara otodidak
- Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
- https://golang.org/pkg/reflect/

```go
type Sample struct {
	Name string
}

type Person struct {
	Name    string
	Address string
	email   string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)

	fmt.Println("Type Name:", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)

		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func main() {

	readField(Sample{"Bani"})
	readField(Person{"Bani", "", ""})
}
```

### Kode Program StructTag

Next: [Package regexp](./15-package-regexp.md)
