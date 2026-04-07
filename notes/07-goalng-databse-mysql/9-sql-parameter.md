# SQL parameter

- Sekarang kita sudah tahu bahaya nya SQL Injection jika menggabungkan string ketika membuat query
- Jika ada kebutuhan seperti itu, sebenarnya function Exec dan Query memiliki parameter tambahan yang bisa kita gunakan untuk mensubtitusi parameter dari function tersebut ke SQL query yang kita buat.
- Untuk menandai sebuah SQL membutuhkan parameter, kita bisa gunakan karakter ? (tanda tanya)

```go
func TestSQLInjectionSafe(t *testing.T) {
	db := GetConnected()
	defer db.Close()

	username := "admin"
	password := "admin"

	ctx := context.Background()
	sqlQuery := "SELECT username FROM users WHERE username=$1 AND password=$2 LIMIT 1;"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses Login", username)

	} else {
		fmt.Println("Gagal Login")
	}

}
```

Next: [auto increment](./10-auto-increment.md)
