# Database Transaction

- Salah satu fitur andalan di database adalah transaction
- Materi database transaction sudah saya bahas dengan tuntas di materi MySQL database, jadi silahkan pelajari di course tersebut
- Di course ini kita akan fokus bagaimana menggunakan database transaction di Golang

## Transaction di Golang

- Secara default, semua perintah SQL yang kita kirim menggunakan Golang akan otomatis di commit, atau istilahnya auto commit
- Namun kita bisa menggunakan fitur transaksi sehingga SQL yang kita kirim tidak secara otomatis di commit ke database
- Untuk memulai transaksi, kita bisa menggunakan function (DB) Begin(), dimana akan menghasilkan struct Tx yang merupakan representasi Transaction
- Struct Tx ini yang kita gunakan sebagai pengganti DB untuk melakukan transaksi, dimana hampir semua function di DB ada di Tx, seperti Exec, Query atau Prepare
- Setelah selesai proses transaksi, kita bisa gunakan function (Tx) Commit() untuk melakukan commit atau Rollback()

```go
func TestTransation(t *testing.T) {

	db := GetConnected()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments (email, comment) VALUES($1, $2) RETURNING id"

	for i := range 10 {
		email := "bani" + strconv.Itoa(i) + "@example.io"
		comment := "Komentar ke:" + strconv.Itoa(i)

		var id int64
		err := tx.QueryRowContext(ctx, script, email, comment).Scan(&id)

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id:", id)
	}

	errorTx := tx.Rollback() // Tidak di commit ke database
	if errorTx != nil {
		panic(errorTx)
	}

}
```

Next: [Repository pattern](./13-repository-pattern.md)
