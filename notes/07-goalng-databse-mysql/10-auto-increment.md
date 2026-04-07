# Auto Increment

- Kadang kita membuat sebuah table dengan id auto increment
- Dan kadang pula, kita ingin mengambil data id yang sudah kita insert ke dalam MySQL
- Sebenarnya kita bisa melakukan query ulang ke database menggunakan SELECT LAST_INSERT_ID()
- Tapi untungnya di Golang ada cara yang lebih mudah
- Kita bisa menggunakan function (Result) LastInsertId() untuk mendapatkan Id terakhir yang dibuat secara auto increment
- Result adalah object yang dikembalikan ketika kita menggunakan function Exec

```go
func TestAutoIncrement(t *testing.T) {
	db := GetConnected()
	defer db.Close()

	ctx := context.Background()

	email := "bani@example.io"
	comment := "Test Komen"

	script := "INSERT INTO comments (email, comment) VALUES($1, $2)"

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	RowsAffected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert new Comment with id", RowsAffected)

}
```

Returning ID Using `QueryRowContext`

```go
func TestAutoIncrement(t *testing.T) {
	db := GetConnected()
	defer db.Close()

	ctx := context.Background()

	email := "bani@example.io"
	comment := "Test Komen"

	script := "INSERT INTO comments (email, comment) VALUES($1, $2) RETURNING id"

	var id int64

	err := db.QueryRowContext(ctx, script, email, comment).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert new Comment with id:", id)

}
```

Next: [Prepare statement](./11-prepare-statement.md)
