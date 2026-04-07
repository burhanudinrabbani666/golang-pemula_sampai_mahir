# Membuka koneksi database

- Hal yang pertama akan kita lakukan ketika membuat aplikasi yang akan menggunakan database adalah melakukan koneksi ke database nya
- Untuk melakukan koneksi ke databsae di Golang, kita bisa membuat object sql.DB menggunakan function sql.Open(driver, dataSourceName)
- Untuk menggunakan database MySQL, kita bisa menggunakan driver “mysql”
- Sedangkan untuk dataSourceName, tiap database biasanya punya cara penulisan masing-masing, misal di MySQL, kita bisa menggunakan dataSourceName seperti dibawah ini :
- username:password@tcp(host:port)/database_name
- Jika object sql.DB sudah tidak digunakan lagi, disarankan untuk menutupnya menggunakan function Close()

```go
import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func TestDatabase(t *testing.T) {}

func TestConnection(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}
	defer db.Close(context.Background())

	// Gunakan DB
	t.Log("Connected!")
}
```

Next: [Database pooling](./4-database-pooling.md)
