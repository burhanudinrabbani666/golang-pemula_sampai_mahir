# Package `path`

## Gambaran Umum

Package `path` digunakan untuk memanipulasi **string path** — seperti path pada URL atau path pada sistem file. Secara default, package ini menggunakan **forward slash (`/`)** sebagai pemisah direktori, sehingga sangat cocok untuk keperluan URL.

> 📖 Dokumentasi resmi: [pkg.go.dev/path](https://pkg.go.dev/path)

---

## `path` vs `path/filepath` — Kapan Pakai Yang Mana?

| Kebutuhan                                 | Package yang Tepat          |
| ----------------------------------------- | --------------------------- |
| Manipulasi path URL                       | `path`                      |
| Manipulasi path File System (Linux/macOS) | `path` atau `path/filepath` |
| Manipulasi path File System (Windows)     | **`path/filepath`**         |

> ⚠️ **Penting:** Windows menggunakan backslash (`\`) sebagai pemisah direktori. Package `path` tidak menangani hal ini — gunakan `path/filepath` agar kode berjalan lintas platform.
>
> 📖 Dokumentasi filepath: [pkg.go.dev/path/filepath](https://pkg.go.dev/path/filepath)

---

## Fungsi-Fungsi Utama

### `path.Dir(p)` — Ambil Direktori

Mengembalikan bagian **direktori** dari sebuah path (semua bagian sebelum nama file terakhir).

### `path.Base(p)` — Ambil Nama File

Mengembalikan bagian **paling akhir** dari sebuah path, biasanya berupa nama file beserta ekstensinya.

### `path.Ext(p)` — Ambil Ekstensi File

Mengembalikan **ekstensi file** (termasuk tanda titik `.`) dari sebuah path.

### `path.Join(elem...)` — Gabungkan Path

Menggabungkan beberapa segmen path menjadi satu path utuh, menggunakan `/` sebagai pemisah. Otomatis membersihkan pemisah berlebih.

---

## Contoh Kode

```go
package main

import (
    "fmt"
    "path"
)

func main() {
    // Mengambil bagian direktori dari path
    fmt.Println(path.Dir("hello/world.go"))             // Output: hello

    // Mengambil nama file (bagian terakhir path)
    fmt.Println(path.Base("hello/world.go"))            // Output: world.go

    // Mengambil ekstensi file
    fmt.Println(path.Ext("hello/world.go"))             // Output: .go

    // Menggabungkan beberapa segmen menjadi satu path
    fmt.Println(path.Join("hello", "world", "main.go")) // Output: hello/world/main.go
}
```

---

## Cara Kerja

```
Input path: "hello/world.go"

  path.Dir()   → Ambil semua SEBELUM "/" terakhir → "hello"
  path.Base()  → Ambil semua SETELAH "/" terakhir  → "world.go"
  path.Ext()   → Ambil dari "." terakhir           → ".go"

path.Join("hello", "world", "main.go")
  → Gabungkan dengan "/"
  → "hello" + "/" + "world" + "/" + "main.go"
  → "hello/world/main.go"
```

---

## Alur Penggunaan Lengkap

```
1. Import package "path" (dan/atau "path/filepath" untuk File System)
        ↓
2. Siapkan string path yang ingin dimanipulasi
        ↓
3. Pilih fungsi sesuai kebutuhan:
   - Dir    → butuh nama folder/direktori
   - Base   → butuh nama file
   - Ext    → butuh ekstensi file
   - Join   → butuh gabungkan beberapa segmen path
        ↓
4. Gunakan hasilnya sebagai string
```

---

## Catatan Penting

| Hal                 | Keterangan                                                      |
| ------------------- | --------------------------------------------------------------- |
| **Pemisah path**    | Selalu menggunakan `/` (forward slash)                          |
| **Lintas platform** | Gunakan `path/filepath` untuk kompatibilitas Windows            |
| **`path.Ext`**      | Mengembalikan string kosong `""` jika file tidak punya ekstensi |
| **`path.Join`**     | Otomatis membersihkan double slash atau trailing slash          |
| **`path.Base`**     | Mengembalikan `"."` jika path adalah string kosong              |
