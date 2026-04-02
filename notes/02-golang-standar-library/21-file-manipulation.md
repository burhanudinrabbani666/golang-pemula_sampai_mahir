# File Manipulation

## Gambaran Umum

Package `os` menyediakan fitur **File Management** untuk membuat, membaca, dan memodifikasi file. Struct `os.File` mengimplementasikan interface `io.Reader` dan `io.Writer`, sehingga kita bisa menggunakan package `io` dan `bufio` untuk operasi baca-tulis yang lebih efisien.

---

## Konsep Penting Sebelum Mulai

```
os.File
  ├── implements io.Reader  → bisa dibaca dengan bufio.NewReader()
  └── implements io.Writer  → bisa ditulis dengan file.WriteString()
```

Karena `os.File` adalah implementasi dari `io.Reader` / `io.Writer`, semua fungsi dari package `io` dan `bufio` dapat langsung digunakan pada file.

---

## Membuka File: `os.OpenFile()`

```go
os.OpenFile(name, flag, permission)
```

| Parameter    | Keterangan                                                                        |
| ------------ | --------------------------------------------------------------------------------- |
| `name`       | Nama file — bisa path absolut (`/home/user/file.txt`) atau relatif (`sample.log`) |
| `flag`       | Mode akses file (lihat tabel flag di bawah)                                       |
| `permission` | Permission Unix saat **membuat** file baru (contoh: `0666`)                       |

### Flag yang Tersedia

| Flag          | Fungsi                                         |
| ------------- | ---------------------------------------------- |
| `os.O_CREATE` | Buat file baru jika belum ada                  |
| `os.O_WRONLY` | Buka hanya untuk menulis                       |
| `os.O_RDONLY` | Buka hanya untuk membaca                       |
| `os.O_RDWR`   | Buka untuk membaca dan menulis                 |
| `os.O_APPEND` | Tambahkan konten di akhir file (tidak menimpa) |

> Flag dapat dikombinasikan menggunakan operator `|`, contoh: `os.O_RDWR|os.O_APPEND`

> 🔑 **Permission:** Gunakan [chmod-calculator.com](https://chmod-calculator.com/) untuk mensimulasikan nilai permission. Nilai `0666` berarti semua user dapat membaca dan menulis.

---

## Contoh Kode

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

// createNewFile membuat file baru dan menulis pesan ke dalamnya.
// Flag O_CREATE membuat file jika belum ada, O_WRONLY membuka hanya untuk menulis.
func createNewFile(name string, message string) error {
    file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        return err
    }
    defer file.Close() // Pastikan file selalu ditutup setelah selesai

    file.WriteString(message)
    return nil
}

// readFile membaca seluruh isi file baris per baris menggunakan bufio.Reader.
// Flag O_RDONLY membuka file hanya untuk dibaca.
func readFile(name string) (string, error) {
    file, err := os.OpenFile(name, os.O_RDONLY, 0666)
    if err != nil {
        return "", err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    var message string
    for {
        line, _, err := reader.ReadLine()

        if err == io.EOF { // EOF (End of File) = tidak ada lagi data yang dibaca
            break
        }

        message += string(line) + "\n"
    }

    return message, nil
}

// addNewFile menambahkan pesan baru di akhir file tanpa menimpa isi sebelumnya.
// Flag O_APPEND memastikan penulisan dilakukan di akhir file.
func addNewFile(name string, message string) error {
    file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    file.WriteString(message)
    return nil
}

func main() {
    // 1. Buat file baru dan tulis konten awal
    createNewFile("sample.log", "This is sample log")

    // 2. Baca dan tampilkan isi file
    result, _ := readFile("sample.log")
    fmt.Println(result)

    // 3. Tambahkan baris baru di akhir file tanpa menghapus isi sebelumnya
    addNewFile("sample.log", "\nThis is additional log")
}
```

---

## Cara Kerja

### Alur `createNewFile`

```
os.OpenFile(O_CREATE | O_WRONLY)
  → Buat file jika belum ada
  → Buka dalam mode tulis saja
  → WriteString() → tulis pesan ke file
  → defer Close() → tutup file otomatis saat fungsi selesai
```

### Alur `readFile`

```
os.OpenFile(O_RDONLY)
  → Buka file dalam mode baca saja
  → bufio.NewReader(file) → bungkus file dengan buffer untuk efisiensi
  → Loop ReadLine():
      ├── Dapat baris → tambahkan ke variabel message
      └── io.EOF      → hentikan loop, kembalikan hasil
```

### Alur `addNewFile`

```
os.OpenFile(O_RDWR | O_APPEND)
  → Buka file untuk baca-tulis
  → O_APPEND → posisi kursor otomatis di akhir file
  → WriteString() → tulis tanpa menimpa konten lama
```

---

## Alur Penggunaan Lengkap

```
1. Tentukan operasi yang dibutuhkan
        ↓
2. Pilih flag yang sesuai:
   - Buat + tulis  → O_CREATE | O_WRONLY
   - Baca saja     → O_RDONLY
   - Tambah isi    → O_RDWR | O_APPEND
        ↓
3. Buka file dengan os.OpenFile()
        ↓
4. Selalu tambahkan defer file.Close()
        ↓
5. Lakukan operasi baca / tulis
        ↓
6. Tangani error yang mungkin terjadi
```

---

## Catatan Penting

| Hal                      | Keterangan                                                                              |
| ------------------------ | --------------------------------------------------------------------------------------- |
| **`defer file.Close()`** | Wajib dipanggil untuk menghindari memory leak / file terkunci                           |
| **`io.EOF`**             | Bukan error nyata — ini sinyal bahwa file sudah habis dibaca                            |
| **`O_APPEND`**           | Tanpa flag ini, penulisan akan menimpa isi file dari awal                               |
| **`bufio.NewReader`**    | Membaca file dengan buffer — lebih efisien untuk file besar                             |
| **Permission `0666`**    | Hanya berlaku saat file **dibuat** (`O_CREATE`), bukan saat membuka file yang sudah ada |

---

Next: [Package Lainnya](./22-package-lainnya.md)
