# Prepare Statement

- Query atau Exec dengan Parameter
- Saat kita menggunakan Function Query atau Exec yang menggunakan parameter, sebenarnya implementasi dibawah nya menggunakan Prepare Statement
- Jadi tahapan pertama statement nya disiapkan terlebih dahulu, setelah itu baru di isi dengan parameter
- Kadang ada kasus kita ingin melakukan beberapa hal yang sama sekaligus, hanya berbeda parameternya. Misal insert data langsung banyak
- Pembuatan Prepare Statement bisa dilakukan dengan manual, tanpa harus mennggunakan Query atau Exec dengan parameter

## Prepare Statement

- Saat kita membuat Prepare Statement, secara otomatis akan mengenali koneksi database yang digunakan
- Sehingga ketika kita mengeksekusi Prepare Statement berkali-kali, maka akan menggunakan koneksi yang sama dan lebih efisien karena pembuatan prepare statement nya hanya sekali diawal saja
- Jika menggunakan Query dan Exec dengan parameter, kita tidak bisa menjamin bahwa koneksi yang digunakan akan sama, oleh karena itu, bisa jadi prepare statement akan selalu dibuat berkali-kali walaupun kita menggunakan SQL yang sama
- Untuk membuat Prepare Statement, kita bisa menggunakan function (DB) Prepare(context, sql)
- Prepare Statement direpresentasikan dalam struct database/sql.Stmt
- Sama seperti resource sql lainnya, Stmt harus di Close() jika sudah tidak digunakan lagi

```go
func TestPrepareStatment(t *testing.T) {
	db := GetConnected()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO comments (email, comment) VALUES($1, $2) RETURNING id"

	statment, err := db.PrepareContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer statment.Close()

	for i := range 10 {
		email := "bani" + strconv.Itoa(i) + "@example.io"
		comment := "Komentar ke:" + strconv.Itoa(i)

		var id int64
		err := statment.QueryRowContext(ctx, email, comment).Scan(&id)

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id:", id)
	}

}
```

Next: [Database transaction](./12-database-transaction.md)
