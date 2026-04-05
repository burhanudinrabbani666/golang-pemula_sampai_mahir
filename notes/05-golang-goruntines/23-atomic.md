# Atomic

## Gambaran Umum

Package **`sync/atomic`** menyediakan operasi pada tipe data primitif (seperti `int32`, `int64`, `bool`) yang **aman digunakan secara concurrent** — tanpa perlu Mutex.

Operasi atomic bekerja di level hardware, memastikan bahwa operasi baca-ubah-tulis terjadi sebagai **satu langkah yang tidak dapat diinterupsi** oleh goroutine lain.

> 📖 Dokumentasi resmi: [pkg.go.dev/sync/atomic](https://pkg.go.dev/sync/atomic)

---

## Atomic vs Mutex — Kapan Pakai Yang Mana?

| Aspek       | `sync.Mutex`                      | `sync/atomic`                           |
| ----------- | --------------------------------- | --------------------------------------- |
| Cocok untuk | Blok kode kompleks, struct        | Operasi sederhana pada satu variabel    |
| Performa    | Lebih lambat (overhead locking)   | Lebih cepat (operasi hardware langsung) |
| Kemudahan   | Lebih fleksibel                   | Hanya untuk tipe primitif               |
| Contoh      | Mengubah beberapa field sekaligus | Increment counter, flag boolean         |

---

## Fungsi-Fungsi Umum

| Fungsi                                     | Fungsi                                                  |
| ------------------------------------------ | ------------------------------------------------------- |
| `atomic.AddInt64(&x, n)`                   | Tambahkan `n` ke `x` secara atomic                      |
| `atomic.LoadInt64(&x)`                     | Baca nilai `x` secara atomic                            |
| `atomic.StoreInt64(&x, n)`                 | Simpan nilai `n` ke `x` secara atomic                   |
| `atomic.CompareAndSwapInt64(&x, old, new)` | Ganti nilai hanya jika nilai saat ini sama dengan `old` |

> Tersedia juga versi untuk `int32`, `uint32`, `uint64`, `uintptr`, dan `unsafe.Pointer`.

---

## Contoh Kode

```go
func TestAtomic(t *testing.T) {
    var x int64 = 0
    group := sync.WaitGroup{}

    // 1000 goroutine berjalan paralel, masing-masing menambah x sebanyak 100 kali
    for i := 1; i <= 1000; i++ {
        group.Go(func() {
            for j := 1; j <= 100; j++ {
                atomic.AddInt64(&x, 1) // Increment atomic — aman tanpa Mutex
            }
        })
    }

    group.Wait()
    fmt.Println("Counter =", x) // Selalu menghasilkan: 100000
}
```

---

## Cara Kerja

```
Tanpa atomic (race condition):
  Goroutine A: baca x=5 → hitung 5+1=6 → tulis x=6
  Goroutine B: baca x=5 → hitung 5+1=6 → tulis x=6  ← lost update!

Dengan atomic.AddInt64:
  Goroutine A: AddInt64(&x, 1) → operasi hardware satu langkah → x=6
  Goroutine B: AddInt64(&x, 1) → menunggu A selesai → x=7  ✅

Tidak ada goroutine yang bisa menginterupsi di tengah operasi atomic.
```

---

## Perbandingan Mutex vs Atomic (Counter)

```go
// Dengan Mutex
mutex.Lock()
x = x + 1
mutex.Unlock()

// Dengan Atomic — lebih ringkas dan lebih cepat untuk kasus ini
atomic.AddInt64(&x, 1)
```

---

## Alur Penggunaan Atomic

```
1. Deklarasikan variabel dengan tipe yang sesuai: var x int64
        ↓
2. Gunakan fungsi atomic untuk setiap akses dari goroutine:
   - Tambah  → atomic.AddInt64(&x, n)
   - Baca    → atomic.LoadInt64(&x)
   - Simpan  → atomic.StoreInt64(&x, n)
        ↓
3. Tidak perlu Lock/Unlock — atomic menangani sinkronisasi secara internal
```

---

## Catatan Penting

| Hal                                     | Keterangan                                                                                       |
| --------------------------------------- | ------------------------------------------------------------------------------------------------ |
| **Hanya untuk primitif**                | Atomic hanya bekerja untuk `int32`, `int64`, `uint32`, dll. — tidak bisa untuk struct atau slice |
| **Selalu gunakan pointer**              | Semua fungsi atomic menerima **pointer** ke variabel (`&x`), bukan nilai langsung                |
| **Lebih cepat dari Mutex**              | Untuk operasi sederhana seperti counter, atomic jauh lebih efisien                               |
| **Tidak menggantikan Mutex sepenuhnya** | Untuk operasi yang melibatkan beberapa variabel sekaligus, Mutex tetap diperlukan                |
| **`group.Go()` tersedia sejak Go 1.22** | Pengganti ringkas dari pola `Add(1)` + `go func()` + `Done()`                                    |

---

Next: [Time Timer](./24-time-timer.md)
