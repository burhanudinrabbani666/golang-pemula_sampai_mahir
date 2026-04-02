# Package time

- Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
- https://golang.org/pkg/time/

## Beberapa Function di Package time

| Function                   | Kegunaan                           |
| -------------------------- | ---------------------------------- |
| time.Now()                 | Untuk mendapatkan waktu saat ini   |
| time.Date(...)             | Untuk membuat waktu                |
| time.Parse(layout, string) | Untuk memparsing waktu dari string |

```go
func main() {

	// Set time now
	var timeNow time.Time = time.Now()
	fmt.Println(timeNow.Local())

	// Set manually
	utc := time.Date(2002, time.November, 14, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	// Using Format
	formater := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formater, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(timeNow.Year())
	fmt.Println(timeNow.Month())
	fmt.Println(timeNow.Date())
	fmt.Println(timeNow.Day())
	fmt.Println(timeNow.Hour())

}
```

## Duration

- Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
- Package tipe memiliki type Duration, yang sebenarnya adalah alias untuk int64
- Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration

```go
	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3)
```

Next: [package reflect](./14-package-reflect.md)
