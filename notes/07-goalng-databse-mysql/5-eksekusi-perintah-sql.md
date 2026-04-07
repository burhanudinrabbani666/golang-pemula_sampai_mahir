# Eksekusi perintah SQL

- Saat membuat aplikasi menggunakan database, sudah pasti kita ingin berkomunikasi dengan database menggunakan perintah SQL
- Di Golang juga menyediakan function yang bisa kita gunakan untuk mengirim perintah SQL ke database menggunakan function `(DB) ExecContext(context, sql, params)`
- Ketika mengirim perintah SQL, kita butuh mengirimkan context, dan seperti yang sudah pernah kita pelajari di course Golang Context, dengan context, kita bisa mengirim sinyal cancel jika kita ingin membatalkan pengiriman perintah SQL nya

```go
func TestExecuteql(t *testing.T) {
	db := GetConnected()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUES('Agus', 'Agus')"

	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Succes Insert new Costomer")

}
```

Next: [query SQL](./6-query-sql.md)
