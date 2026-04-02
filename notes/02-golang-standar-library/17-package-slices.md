# Package `slices`

## Gambaran Umum

Package `slices` adalah bagian dari standard library Go yang memanfaatkan fitur **Generics** (diperkenalkan di Go 1.21). Package ini menyediakan fungsi-fungsi utilitas untuk memanipulasi data dalam slice secara type-safe — tanpa harus menggunakan `interface{}` / `any`.

> 📖 Dokumentasi resmi: [pkg.go.dev/slices](https://pkg.go.dev/slices)

---

## Prasyarat: Memahami Generics

Sebelum menggunakan package ini, penting untuk mengetahui:

- **Generics** memungkinkan sebuah fungsi menerima parameter dengan tipe data yang fleksibel (bisa `string`, `int`, dll.)
- Tanpa generics, kita harus membuat fungsi terpisah untuk setiap tipe, atau menggunakan `any` yang tidak type-safe
- Package `slices` menggunakan generics sehingga satu fungsi seperti `slices.Min()` dapat bekerja pada `[]string`, `[]int`, maupun tipe lain yang dapat dibandingkan

---

## Fungsi-Fungsi Utama

### `slices.Min(s)` — Nilai Terkecil

Mengembalikan elemen terkecil dalam slice. Untuk `[]string`, perbandingan dilakukan secara leksikografis (urutan alfabet).

### `slices.Max(s)` — Nilai Terbesar

Mengembalikan elemen terbesar dalam slice. Cara kerja sama dengan `Min`, namun hasilnya kebalikan.

### `slices.Contains(s, v)` — Cek Keberadaan Elemen

Mengembalikan `true` jika nilai `v` ditemukan dalam slice `s`, dan `false` jika tidak ada.

### `slices.Index(s, v)` — Posisi Elemen

Mengembalikan **indeks pertama** dari nilai `v` dalam slice `s`. Mengembalikan `-1` jika nilai tidak ditemukan.

---

## Contoh Kode

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    // Slice bertipe string dan int
    names := []string{"Jhon", "Paul", "George", "Ringo"}
    values := []int{100, 95, 90, 80}

    // Mencari nilai terkecil
    fmt.Println(slices.Min(names))   // Output: "George" (alfabet terkecil)
    fmt.Println(slices.Min(values))  // Output: 80

    // Mencari nilai terbesar
    fmt.Println(slices.Max(names))   // Output: "Ringo" (alfabet terbesar)
    fmt.Println(slices.Max(values))  // Output: 100

    // Mengecek apakah elemen ada dalam slice
    fmt.Println(slices.Contains(names, "Bani"))  // Output: false (tidak ada)

    // Mencari indeks elemen
    fmt.Println(slices.Index(names, "Bani"))  // Output: -1 (tidak ditemukan)
    fmt.Println(slices.Index(names, "Paul"))  // Output: 1 (indeks ke-1)
}
```

---

## Cara Kerja

```
names := ["Jhon", "Paul", "George", "Ringo"]
indeks:      0       1        2         3

slices.Min(names)
  → Membandingkan semua elemen secara leksikografis
  → "George" < "Jhon" < "Paul" < "Ringo"
  → Hasil: "George"

slices.Index(names, "Paul")
  → Iterasi dari indeks 0, 1, 2 ...
  → Ditemukan di indeks 1
  → Hasil: 1

slices.Index(names, "Bani")
  → Iterasi seluruh slice, tidak ditemukan
  → Hasil: -1
```

---

## Alur Penggunaan Lengkap

```
1. Import package "slices"
        ↓
2. Siapkan slice bertipe apapun (string, int, float, dsb.)
        ↓
3. Panggil fungsi yang dibutuhkan:
   - Min / Max  → untuk mencari batas nilai
   - Contains   → untuk validasi keberadaan elemen
   - Index      → untuk mengetahui posisi elemen
        ↓
4. Gunakan hasilnya (nilai, bool, atau indeks)
```

---

## Catatan Penting

| Hal                               | Keterangan                                                |
| --------------------------------- | --------------------------------------------------------- |
| **Versi Go**                      | Tersedia mulai Go 1.21                                    |
| **Perbandingan string**           | Bersifat leksikografis (alfabet), bukan panjang string    |
| **Nilai `Index` tidak ditemukan** | Mengembalikan `-1`, bukan error                           |
| **Type-safe**                     | Compiler akan menolak tipe yang tidak kompatibel          |
| **Panic**                         | `Min` dan `Max` akan panic jika slice kosong (`len == 0`) |

---

Next: [Package path](./18-package-path.md)
