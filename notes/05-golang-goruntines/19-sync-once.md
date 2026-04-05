# Sync Once

## Gambaran Umum

**`sync.Once`** menjamin bahwa sebuah fungsi hanya dieksekusi **tepat satu kali**, tidak peduli berapa banyak goroutine yang mencoba memanggilnya. Goroutine pertama yang berhasil akan menjalankan fungsinya — goroutine lain yang datang belakangan akan diabaikan.

Ini sangat berguna untuk keperluan **inisialisasi satu kali** seperti membuat koneksi database, memuat konfigurasi, atau membuat singleton.

---

## Method Once

| Method       | Fungsi                                                                                          |
| ------------ | ----------------------------------------------------------------------------------------------- |
| `once.Do(f)` | Menjalankan fungsi `f` **hanya sekali** — panggilan berikutnya tidak akan mengeksekusi `f` lagi |

---

## Contoh Kode

```go
var counter = 0

// OnlyOnce adalah fungsi yang kita inginkan hanya dieksekusi satu kali
// meskipun dipanggil dari 100 goroutine
func OnlyOnce() {
    counter++
}

func TestOnce(t *testing.T) {
    once  := sync.Once{}
    group := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        // group.Go() adalah cara ringkas pengganti Add(1) + go func() + Done()
        group.Go(func() {
            once.Do(OnlyOnce) // Hanya goroutine PERTAMA yang berhasil mengeksekusi OnlyOnce()
                              // 99 goroutine lainnya memanggil Do() tapi tidak melakukan apa-apa
        })
    }

    group.Wait()
    fmt.Println("Counter:", counter) // Output: 1 — bukan 100
}
```

**Output:**

```
Counter: 1
```

---

## Cara Kerja

```
100 goroutine memanggil once.Do(OnlyOnce) secara bersamaan

  Goroutine 1  → once.Do() → ✅ eksekusi OnlyOnce() → counter = 1
  Goroutine 2  → once.Do() → ❌ sudah dieksekusi, diabaikan
  Goroutine 3  → once.Do() → ❌ sudah dieksekusi, diabaikan
  ...
  Goroutine 100 → once.Do() → ❌ sudah dieksekusi, diabaikan

Hasil akhir: counter = 1 (selalu, tidak peduli urutan goroutine)
```

---

## `group.Go()` vs Cara Manual

Kode di atas menggunakan `group.Go()` sebagai cara ringkas. Berikut perbandingannya:

```go
// ✅ Cara ringkas dengan group.Go() — tersedia di Go 1.22+
group.Go(func() {
    once.Do(OnlyOnce)
})

// ✅ Cara manual yang setara
group.Add(1)
go func() {
    once.Do(OnlyOnce)
    group.Done()
}()
```

Keduanya menghasilkan perilaku yang sama — `group.Go()` hanya lebih singkat karena menggabungkan `Add(1)`, `go`, dan `Done()` dalam satu panggilan.

---

## Alur Penggunaan Sync Once

```
1. Deklarasikan: once := sync.Once{}
        ↓
2. Bungkus fungsi yang ingin dijalankan sekali dengan once.Do()
        ↓
3. Panggil once.Do() dari goroutine manapun — hanya eksekusi pertama yang berlaku
        ↓
4. Semua panggilan berikutnya ke once.Do() diabaikan secara otomatis
```

---

## Catatan Penting

| Hal                       | Keterangan                                                                 |
| ------------------------- | -------------------------------------------------------------------------- |
| **Eksekusi tepat sekali** | Dijamin hanya satu kali meski dipanggil dari ribuan goroutine              |
| **Thread-safe**           | `sync.Once` aman digunakan dari banyak goroutine secara bersamaan          |
| **Tidak bisa di-reset**   | Setelah `Do()` dijalankan sekali, tidak ada cara untuk menjalankannya lagi |
| **`group.Go()`**          | Tersedia sejak Go 1.22, menggantikan pola `Add(1)` + `go` + `Done()`       |
| **Kasus penggunaan**      | Inisialisasi singleton, koneksi database, load konfigurasi                 |

---

Next: [Sync Pool](./20-sync-pool.md)
