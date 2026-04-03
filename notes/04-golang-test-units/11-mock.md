# Mock

- Mock adalah object yang sudah kita program dengan ekspektasi tertentu sehingga ketika dipanggil, dia akan menghasilkan data yang sudah kita program diawal
- Mock adalah salah satu teknik dalam unit testing, dimana kita bisa membuat mock object dari suatu object yang memang sulit untuk di testing
- Misal kita ingin membuat unit test, namun ternyata ada kode program kita yang harus memanggil API Call ke third party service. Hal ini sangat sulit untuk di test, karena unit testing kita harus selalu memanggil third party service, dan belum tentu response nya sesuai dengan apa yang kita mau
- Pada kasus seperti ini, cocok sekali untuk menggunakan mock object

## Testify Mock

- Untuk membuat mock object, tidak ada fitur bawaan Go-Lang, namun kita bisa menggunakan library testify yang sebelumnya kita gunakan untuk assertion
- Testify mendukung pembuatan mock object, sehingga cocok untuk kita gunakan ketika ingin membuat mock object
- Namun, perlu diperhatikan, jika desain kode program kita jelek, akan sulit untuk melakukan mocking, jadi pastikan kita melakukan pembuatan desain kode program kita dengan baik
- Mari kita buat contoh kasus

## Aplikasi Query Ke Database

- Kita akan coba contoh kasus dengan membuat contoh aplikasi golang yang melakukan query ke database
- Dimana kita akan buat layer Service sebagai business logic, dan layer Repository sebagai jembatan ke database
- Agar kode kita mudah untuk di test, disarankan agar membuat kontrak berupa Interface

## Gambaran Umum

**Mock** adalah object yang sudah diprogram dengan ekspektasi tertentu — ketika dipanggil, ia akan mengembalikan data yang sudah kita tentukan sebelumnya.

Mock digunakan dalam **unit testing** untuk menggantikan bagian kode yang sulit diuji secara langsung, seperti:

- Pemanggilan API ke third-party service
- Query ke database
- Operasi file atau jaringan

> **Mengapa perlu Mock?** Bayangkan unit test yang harus memanggil API eksternal setiap kali dijalankan — lambat, tidak stabil, dan hasilnya tidak bisa dikontrol. Dengan mock, kita **simulasikan** responnya sehingga test tetap cepat, terisolasi, dan deterministik.

---

## Testify Mock

Go tidak menyediakan fitur mock bawaan, namun library **Testify** mendukung pembuatan mock object melalui package `github.com/stretchr/testify/mock`.

> ⚠️ **Perhatian:** Mock hanya bekerja baik jika desain kode menggunakan **Interface**. Tanpa interface, mock tidak dapat diterapkan. Pastikan arsitektur kode sudah menggunakan kontrak interface sebelum mencoba mocking.

---

## Studi Kasus: Aplikasi Query ke Database

Kita akan membangun simulasi aplikasi dengan dua layer:

```
[ CategoryService ]  ← business logic
        ↓
[ CategoryRepository ]  ← interface sebagai kontrak
        ↓
[ Database / Mock ]  ← implementasi nyata atau mock
```

Dengan struktur ini, `CategoryService` tidak bergantung langsung pada database — melainkan pada **interface** `CategoryRepository`. Saat testing, kita bisa menukar implementasi database dengan mock.

---

### 1. Entity — Struktur Data

```go
// File: entity/category.go
package entity

// Category merepresentasikan data kategori dari database
type Category struct {
    Id   string
    Name string
}
```

---

### 2. Repository Interface — Kontrak Akses Data

```go
// File: repository/category_repository.go
package repository

import "golang-pemula_sampai_mahir/entity"

// CategoryRepository adalah kontrak (interface) yang harus diimplementasikan
// oleh siapapun yang ingin menjadi sumber data Category — baik database nyata maupun mock
type CategoryRepository interface {
    FindById(id string) *entity.Category
}
```

---

### 3. Mock Repository — Implementasi Palsu untuk Testing

```go
// File: repository/category_repository_mock.go
package repository

import (
    "golang-pemula_sampai_mahir/entity"
    "github.com/stretchr/testify/mock"
)

// CategoryRepositoryMock mengimplementasikan interface CategoryRepository
// namun alih-alih query ke database, ia mengembalikan data yang sudah diprogram
type CategoryRepositoryMock struct {
    Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
    // Beritahu testify bahwa fungsi ini dipanggil dengan argumen id
    arguments := repository.Mock.Called(id)

    if arguments.Get(0) == nil {
        return nil // Simulasi: data tidak ditemukan
    } else {
        // Type assertion: ubah interface{} menjadi entity.Category
        category := arguments.Get(0).(entity.Category)
        return &category
    }
}
```

---

### 4. Service — Business Logic

```go
// File: service/category_service.go
package service

import (
    "errors"
    "golang-pemula_sampai_mahir/entity"
    "golang-pemula_sampai_mahir/repository"
)

// CategoryService berisi business logic yang bergantung pada repository
// Service ini tidak tahu apakah repository-nya database nyata atau mock
type CategoryService struct {
    Repository repository.CategoryRepository
}

// Get mencari category berdasarkan id, mengembalikan error jika tidak ditemukan
func (service CategoryService) Get(id string) (*entity.Category, error) {
    category := service.Repository.FindById(id)

    if category == nil {
        return nil, errors.New("Category Not Found")
    }

    return category, nil
}
```

---

### 5. Unit Test — Menggunakan Mock

```go
// File: service/category_service_test.go
package service

import (
    "golang-pemula_sampai_mahir/entity"
    "golang-pemula_sampai_mahir/repository"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// Inisialisasi mock repository dan service yang menggunakannya
var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

// TestCategoryService_GetNotFound memastikan error dikembalikan jika data tidak ada
func TestCategoryService_GetNotFound(t *testing.T) {
    // Program mock: jika FindById dipanggil dengan "1", kembalikan nil
    categoryRepository.Mock.On("FindById", "1").Return(nil)

    category, err := categoryService.Get("1")

    assert.Nil(t, category)    // category harus nil
    assert.NotNil(t, err)      // error harus ada
}

// TestCategoryService_GetSuccess memastikan data dikembalikan jika ditemukan
func TestCategoryService_GetSuccess(t *testing.T) {
    category := entity.Category{
        Id:   "2",
        Name: "Laptop",
    }

    // Program mock: jika FindById dipanggil dengan "2", kembalikan category di atas
    categoryRepository.Mock.On("FindById", "2").Return(category)

    result, err := categoryService.Get("2")

    assert.Nil(t, err)                       // tidak boleh ada error
    assert.NotNil(t, result)                 // result harus ada
    assert.Equal(t, category.Id, result.Id)  // Id harus sama
    assert.Equal(t, category.Name, result.Name) // Name harus sama
}
```

---

## Cara Kerja Mock

```
Test memanggil categoryService.Get("1")
        ↓
Service memanggil repository.FindById("1")
        ↓
Mock mencocokkan argumen "1" dengan ekspektasi yang sudah diprogram
        ↓
Mock.On("FindById", "1").Return(nil)  → kembalikan nil
        ↓
Service menerima nil → return error "Category Not Found"
        ↓
Test memverifikasi: category == nil ✓, err != nil ✓
```

---

## Alur Penggunaan Mock

```
1. Buat Interface sebagai kontrak repository
        ↓
2. Buat struct Mock yang mengimplementasikan interface tersebut
        ↓
3. Gunakan mock.Mock.Called() di setiap method mock
        ↓
4. Di test, program ekspektasi dengan Mock.On("method", arg).Return(nilai)
        ↓
5. Jalankan fungsi yang ingin ditest
        ↓
6. Verifikasi hasil dengan assert
```

---

## Catatan Penting

| Hal                 | Keterangan                                                                                |
| ------------------- | ----------------------------------------------------------------------------------------- |
| **Interface wajib** | Mock hanya bisa digunakan jika kode menggunakan interface                                 |
| **`Mock.On()`**     | Mendefinisikan ekspektasi: "jika method ini dipanggil dengan argumen ini, kembalikan ini" |
| **`Mock.Called()`** | Dipanggil di dalam method mock untuk mencatat bahwa method telah dipanggil                |
| **Type assertion**  | `arguments.Get(0).(entity.Category)` mengubah `interface{}` ke tipe konkret               |
| **Isolasi test**    | Mock memastikan test tidak bergantung pada database atau service eksternal                |

---

Next: [Benchmark](./12-benchmark.md)
