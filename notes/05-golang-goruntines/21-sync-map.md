# Sync Map

## Gambaran Umum

**`sync.Map`** adalah versi map bawaan Go yang **aman digunakan secara concurrent** dari banyak goroutine sekaligus. Berbeda dengan map biasa (`map[K]V`) yang akan menyebabkan race condition jika diakses dari beberapa goroutine secara bersamaan, `sync.Map` sudah dilengkapi mekanisme sinkronisasi internal.

---

## Map Biasa vs sync.Map

| Aspek                     | `map[K]V`                 | `sync.Map`            |
| ------------------------- | ------------------------- | --------------------- |
| Akses concurrent          | ❌ Race condition         | ✅ Thread-safe        |
| Sintaks                   | `m[key] = value`          | `m.Store(key, value)` |
| Tipe key/value            | Ditentukan saat deklarasi | `any` (interface{})   |
| Performa single goroutine | Lebih cepat               | Sedikit lebih lambat  |
| Cocok untuk               | Satu goroutine            | Banyak goroutine      |

---

## Method sync.Map

| Method                               | Fungsi                                                                        |
| ------------------------------------ | ----------------------------------------------------------------------------- |
| `m.Store(key, value)`                | Menyimpan pasangan key-value ke dalam map                                     |
| `m.Load(key)`                        | Mengambil nilai berdasarkan key — mengembalikan `(value, ok bool)`            |
| `m.Delete(key)`                      | Menghapus entri berdasarkan key                                               |
| `m.Range(func(key, value any) bool)` | Iterasi seluruh data — kembalikan `true` untuk lanjut, `false` untuk berhenti |

---

## Contoh Kode

```go
package golanggoroutine

import (
    "fmt"
    "sync"
    "testing"
)

// AddToMap menyimpan satu nilai ke sync.Map secara concurrent
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
    defer group.Done() // Kurangi counter WaitGroup saat selesai

    data.Store(value, value) // Simpan key=value, val=value ke map
}

func TestMap(t *testing.T) {
    data  := &sync.Map{}
    group := &sync.WaitGroup{}

    // Jalankan 100 goroutine yang masing-masing menyimpan satu nilai ke map
    for i := 0; i < 100; i++ {
        group.Add(1)
        go AddToMap(data, i, group)
    }

    group.Wait() // Tunggu semua goroutine selesai menyimpan data

    // Iterasi seluruh isi map setelah semua goroutine selesai
    data.Range(func(key, value any) bool {
        fmt.Println(key, ":", value)
        return true // Kembalikan true untuk melanjutkan iterasi ke entri berikutnya
    })
}
```

---

## Cara Kerja

```
100 goroutine berjalan paralel:
  Goroutine 0  → data.Store(0, 0)  ✅
  Goroutine 1  → data.Store(1, 1)  ✅
  ...
  Goroutine 99 → data.Store(99, 99) ✅

group.Wait() → menunggu semua goroutine selesai
        ↓
data.Range() → iterasi semua 100 entri
  key: 0, value: 0
  key: 1, value: 1
  ...  (urutan tidak dijamin)
  key: 99, value: 99
```

---

## Penggunaan `Load` dan `Delete`

```go
// Menyimpan data
data.Store("nama", "Bani")

// Mengambil data
value, ok := data.Load("nama")
if ok {
    fmt.Println(value) // Output: Bani
}

// Menghapus data
data.Delete("nama")

// Cek setelah dihapus
_, ok = data.Load("nama")
fmt.Println(ok) // Output: false
```

---

## Alur Penggunaan Sync Map

```
1. Buat sync.Map: data := &sync.Map{}
        ↓
2. Simpan data dari goroutine: data.Store(key, value)
        ↓
3. Baca data: value, ok := data.Load(key)
        ↓
4. Hapus data: data.Delete(key)
        ↓
5. Iterasi semua data: data.Range(func(k, v any) bool { return true })
```

---

## Catatan Penting

| Hal                                     | Keterangan                                                                     |
| --------------------------------------- | ------------------------------------------------------------------------------ |
| **Thread-safe**                         | Tidak perlu Mutex tambahan — `sync.Map` sudah menangani sinkronisasi internal  |
| **Tipe `any`**                          | Key dan value bertipe `any` — perlu type assertion saat menggunakan nilainya   |
| **Urutan `Range` tidak dijamin**        | Iterasi tidak mengikuti urutan penyisipan — urutan acak setiap kali dijalankan |
| **`Load` mengembalikan dua nilai**      | Selalu cek nilai `ok` untuk memastikan key memang ada di map                   |
| **Jangan pakai untuk single goroutine** | Map biasa (`map[K]V`) lebih cepat jika tidak ada concurrent access             |

---

Next: [Sync Cond](./22-sync-cond.md)
