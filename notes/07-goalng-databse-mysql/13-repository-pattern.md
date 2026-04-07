# Repository Pattern

## Gambaran Umum

**Repository Pattern** adalah pola desain yang memisahkan **business logic** dari **logika akses data** (SQL, database). Semua perintah SQL dikumpulkan di dalam Repository, sehingga kode bisnis tidak perlu tahu detail implementasi database.

> Menurut Eric Evans dalam buku _Domain-Driven Design_:
> "Repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of objects."

---

## Mengapa Perlu Repository Pattern?

Tanpa pattern ini, kode SQL tersebar di seluruh aplikasi — sulit dirawat, sulit ditest, dan sulit diganti implementasinya.

```
Tanpa Repository:
  Service → SQL langsung → Database
  (SQL tersebar, sulit ditest, susah diganti)

Dengan Repository:
  Service → Repository Interface → Implementasi → Database
  (SQL terpusat, mudah ditest dengan mock, mudah diganti)
```

---

## Struktur Proyek

```
project/
├── entity/
│   └── comment.go          ← Struct representasi tabel database
├── repository/
│   ├── comment_repository.go       ← Interface (kontrak)
│   ├── comment_repository_impl.go  ← Implementasi konkret
│   └── comment_repository_test.go  ← Unit test
└── db.go                   ← Koneksi database
```

---

## Langkah 1: Entity — Representasi Tabel Database

Di Go, tabel database direpresentasikan sebagai **struct**. Setiap field struct merepresentasikan satu kolom tabel.

```go
// File: entity/comment.go
package entity

// Comment merepresentasikan satu baris data dari tabel `comments` di database.
// Struct ini digunakan untuk membawa data antar layer (repository → service → handler).
type Comment struct {
    Id      int32  // Kolom: id (primary key, auto-increment)
    Email   string // Kolom: email
    Comment string // Kolom: comment
}
```

**Mengapa menggunakan struct, bukan map atau array?**

- Tipe data jelas dan aman (type-safe)
- Mudah dibaca dan dipahami
- IDE bisa memberikan autocomplete dan deteksi error

---

## Langkah 2: Interface — Kontrak Repository

Interface mendefinisikan **operasi apa saja** yang bisa dilakukan pada data Comment. Ini adalah "kontrak" yang harus dipenuhi oleh setiap implementasi.

```go
// File: repository/comment_repository.go
package repository

import (
    "context"
    "golang-database-postgre/entity"
)

// CommentRepository mendefinisikan kontrak akses data untuk entitas Comment.
// Siapapun yang mengimplementasikan interface ini bisa digunakan sebagai
// repository — baik implementasi nyata (database) maupun mock (untuk testing).
type CommentRepository interface {
    // Insert menyimpan comment baru ke database.
    // Mengembalikan comment lengkap dengan Id yang dihasilkan database.
    Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)

    // FindById mencari satu comment berdasarkan Id.
    // Mengembalikan error jika Id tidak ditemukan.
    FindById(ctx context.Context, id int32) (entity.Comment, error)

    // FindAll mengambil semua comment dari database.
    // Mengembalikan slice kosong (bukan nil) jika tidak ada data.
    FindAll(ctx context.Context) ([]entity.Comment, error)
}
```

**Mengapa menggunakan Interface?**

- Memungkinkan **mock** saat testing (tidak perlu database nyata)
- Memudahkan **penggantian implementasi** (dari PostgreSQL ke MySQL misalnya) tanpa mengubah kode bisnis
- Menerapkan prinsip **Dependency Injection**

---

## Langkah 3: Implementasi Repository

Implementasi konkret yang berisi logika SQL aktual. Struct ini tidak diekspor (`commentRepositoryImpl`) — hanya bisa diakses melalui constructor `NewCommentRepository`.

```go
// File: repository/comment_repository_impl.go
package repository

import (
    "context"
    "database/sql"
    "errors"
    "golang-database-postgre/entity"
    "strconv"
)

// commentRepositoryImpl adalah implementasi konkret dari CommentRepository.
// Menggunakan huruf kecil (unexported) agar hanya bisa dibuat melalui constructor.
type commentRepositoryImpl struct {
    DB *sql.DB // Koneksi database yang diinjeksikan dari luar
}

// NewCommentRepository adalah constructor untuk membuat instance repository.
// Menerima *sql.DB dari luar (Dependency Injection) agar mudah ditest.
// Mengembalikan interface, bukan struct konkret — menyembunyikan detail implementasi.
func NewCommentRepository(db *sql.DB) CommentRepository {
    return &commentRepositoryImpl{DB: db}
}

// Insert menyimpan comment baru dan mengembalikan comment beserta Id yang digenerate database.
// Menggunakan RETURNING id untuk langsung mendapatkan Id tanpa query tambahan.
func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
    // $1, $2 adalah placeholder untuk PostgreSQL (MySQL menggunakan ?)
    // RETURNING id mengambil Id yang dihasilkan database setelah INSERT
    script := "INSERT INTO comments(email, comment) VALUES ($1, $2) RETURNING id"
    var id int64

    err := repository.DB.QueryRowContext(ctx, script, comment.Email, comment.Comment).Scan(&id)
    if err != nil {
        panic(err) // Panic karena error INSERT biasanya tidak bisa di-recover
    }

    comment.Id = int32(id) // Set Id ke struct sebelum dikembalikan
    return comment, nil
}

// FindById mencari satu comment berdasarkan Id menggunakan LIMIT 1.
// Mengembalikan error dengan pesan deskriptif jika Id tidak ditemukan.
func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
    script := "SELECT id, email, comment FROM comments WHERE id = $1 LIMIT 1;"

    rows, err := repository.DB.QueryContext(ctx, script, id)
    if err != nil {
        panic(err)
    }
    defer rows.Close() // Pastikan rows selalu ditutup untuk menghindari connection leak

    comment := entity.Comment{}

    if rows.Next() {
        // Data ditemukan — scan kolom ke field struct
        rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
        return comment, nil
    }

    // Data tidak ditemukan — kembalikan error deskriptif
    return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
}

// FindAll mengambil semua comment dari database.
// Menggunakan loop rows.Next() untuk membaca setiap baris hasil query.
func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
    script := "SELECT id, email, comment FROM comments;"

    rows, err := repository.DB.QueryContext(ctx, script)
    if err != nil {
        return nil, err // Kembalikan nil + error jika query gagal
    }
    defer rows.Close()

    var comments []entity.Comment

    for rows.Next() {
        comment := entity.Comment{}
        rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
        comments = append(comments, comment)
    }

    return comments, nil
}
```

---

## Langkah 4: Testing

```go
// File: repository/comment_repository_test.go
package repository

import (
    "context"
    "fmt"
    golangdatabasepostgre "golang-database-postgre"
    "golang-database-postgre/entity"
    "testing"

    _ "github.com/lib/pq" // Driver PostgreSQL — diimport untuk side effect registrasi driver
)

// TestCommentInsert menguji Insert — menyimpan comment baru ke database
func TestCommentInsert(t *testing.T) {
    commentRepository := NewCommentRepository(golangdatabasepostgre.GetConnected())
    ctx := context.Background()

    comment := entity.Comment{
        Email:   "bani@gmail.com",
        Comment: "Test Repository",
    }

    result, err := commentRepository.Insert(ctx, comment)
    if err != nil {
        panic(err.Error())
    }

    fmt.Println(result) // Output: {Id: <auto>, Email: "bani@gmail.com", Comment: "Test Repository"}
}

// TestFindById menguji FindById — mengambil satu comment berdasarkan Id
func TestFindById(t *testing.T) {
    commentRepository := NewCommentRepository(golangdatabasepostgre.GetConnected())
    ctx := context.Background()

    comment, err := commentRepository.FindById(ctx, 90)
    if err != nil {
        panic(err.Error()) // Akan panic jika Id 90 tidak ada di database
    }

    fmt.Println(comment)
}

// TestFindAll menguji FindAll — mengambil seluruh data comment
func TestFindAll(t *testing.T) {
    commentRepository := NewCommentRepository(golangdatabasepostgre.GetConnected())
    ctx := context.Background()

    comments, err := commentRepository.FindAll(ctx)
    if err != nil {
        panic(err.Error())
    }

    for _, comment := range comments {
        fmt.Println(comment)
    }
}
```

---

## Cara Kerja Lengkap

```
[ Test / Service ]
       │
       │ Memanggil method melalui Interface
       ▼
[ CommentRepository Interface ]
   Insert() / FindById() / FindAll()
       │
       │ Diarahkan ke implementasi konkret
       ▼
[ commentRepositoryImpl ]
       │
       │ Menjalankan SQL query
       ▼
[ *sql.DB ]
       │
       │ Mengirim query ke database
       ▼
[ PostgreSQL Database ]
       │
       │ Mengembalikan hasil
       ▼
[ Rows / Data ]
       │
       │ Di-scan ke struct entity.Comment
       ▼
[ entity.Comment / []entity.Comment ]
       │
       │ Dikembalikan ke pemanggil
       ▼
[ Test / Service ]
```

---

## Alur Penggunaan Repository Pattern

```
1. Buat struct Entity yang merepresentasikan tabel database
        ↓
2. Buat Interface yang mendefinisikan operasi yang tersedia
        ↓
3. Buat Implementasi konkret yang mengisi logika SQL
        ↓
4. Gunakan Constructor (NewXxxRepository) untuk membuat instance
        ↓
5. Gunakan Interface (bukan struct konkret) di service/handler
        ↓
6. Saat testing, ganti implementasi nyata dengan Mock
```

---

## Catatan Penting

| Hal                                     | Keterangan                                                                                                             |
| --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| **`defer rows.Close()`**                | Wajib dipanggil setelah `QueryContext` untuk menghindari connection leak                                               |
| **`$1, $2` vs `?`**                     | PostgreSQL menggunakan `$1`, `$2`; MySQL menggunakan `?` sebagai placeholder                                           |
| **Constructor mengembalikan Interface** | `NewCommentRepository` mengembalikan `CommentRepository`, bukan `*commentRepositoryImpl` — menyembunyikan implementasi |
| **`RETURNING id`**                      | Cara PostgreSQL untuk mendapatkan nilai kolom setelah INSERT tanpa query tambahan                                      |
| **`context.Context` di setiap method**  | Memungkinkan pembatalan query menggunakan `WithCancel` atau `WithTimeout`                                              |
| **`panic` vs `return error`**           | Gunakan `panic` hanya untuk error yang tidak bisa di-recover; untuk error normal, kembalikan sebagai nilai return      |
| **Typo di kode asli**                   | `TestFindAl` diperbaiki menjadi `TestFindAll`; `"bani@gmail.hi"` diperbaiki menjadi `"bani@gmail.com"`                 |
