# Sync Cond

## Gambaran Umum

**`sync.Cond`** adalah mekanisme sinkronisasi berbasis **kondisi** — memungkinkan goroutine untuk **menunggu** sampai ada sinyal dari goroutine lain bahwa kondisi tertentu sudah terpenuhi.

Berbeda dengan `Mutex` yang hanya mengatur akses eksklusif, `Cond` memungkinkan goroutine untuk **tidur** (tidak memakan CPU) sampai dibangunkan melalui sinyal.

---

## Method sync.Cond

| Method                 | Fungsi                                                                                                |
| ---------------------- | ----------------------------------------------------------------------------------------------------- |
| `cond.Wait()`          | Goroutine berhenti menunggu, melepas lock sementara — akan dibangunkan oleh `Signal` atau `Broadcast` |
| `cond.Signal()`        | Membangunkan **satu** goroutine yang sedang `Wait()`                                                  |
| `cond.Broadcast()`     | Membangunkan **semua** goroutine yang sedang `Wait()`                                                 |
| `sync.NewCond(locker)` | Membuat Cond baru — membutuhkan `Locker` (`*sync.Mutex` atau `*sync.RWMutex`)                         |

---

## Contoh Kode

```go
var locker = sync.Mutex{}
var cond   = sync.NewCond(&locker) // Cond membutuhkan Locker sebagai dasarnya
var group  = sync.WaitGroup{}

// WaitCondition adalah goroutine yang menunggu sinyal sebelum melanjutkan
func WaitCondition(value int) {
    defer group.Done()

    cond.L.Lock()   // Kunci locker sebelum memanggil Wait
    cond.Wait()     // Lepas lock sementara dan tidur — menunggu Signal() atau Broadcast()
                    // Setelah dibangunkan, lock diambil kembali secara otomatis
    fmt.Println("Done", value)
    cond.L.Unlock() // Lepas lock setelah selesai
}

func TestCond(t *testing.T) {
    for i := 0; i < 10; i++ {
        group.Add(1)
        go WaitCondition(i) // Jalankan 10 goroutine — semuanya akan menunggu sinyal
    }

    // Goroutine pengirim sinyal — membangunkan satu goroutine per detik
    go func() {
        for i := 0; i < 10; i++ {
            time.Sleep(1 * time.Second)
            cond.Signal() // Bangunkan satu goroutine yang sedang Wait()
        }
    }()

    group.Wait() // Tunggu semua 10 goroutine selesai
}
```

---

## Cara Kerja

```
10 goroutine dibuat → semuanya memanggil cond.Wait() → semua tidur

  Goroutine 0  → Lock → Wait() → 💤 tidur
  Goroutine 1  → Lock → Wait() → 💤 tidur
  ...
  Goroutine 9  → Lock → Wait() → 💤 tidur

Setiap 1 detik, goroutine pengirim memanggil cond.Signal():
  Detik 1: Signal() → bangunkan Goroutine 0 → "Done 0" → Done()
  Detik 2: Signal() → bangunkan Goroutine 1 → "Done 1" → Done()
  ...
  Detik 10: Signal() → bangunkan Goroutine 9 → "Done 9" → Done()

group.Wait() selesai setelah semua 10 goroutine memanggil Done()
```

---

## Signal vs Broadcast

```go
cond.Signal()    // Bangunkan SATU goroutine — mana yang dibangunkan tidak bisa diprediksi
cond.Broadcast() // Bangunkan SEMUA goroutine yang sedang Wait() sekaligus
```

**Gunakan `Broadcast`** jika kondisi yang ditunggu berlaku untuk semua goroutine sekaligus, misalnya "data sudah siap" atau "server sudah berjalan".

---

## Alur Penggunaan Sync Cond

```
1. Buat Locker: var mu sync.Mutex
        ↓
2. Buat Cond: cond := sync.NewCond(&mu)
        ↓
3. Di goroutine yang menunggu:
   cond.L.Lock() → cond.Wait() → (lanjut setelah dibangunkan) → cond.L.Unlock()
        ↓
4. Di goroutine pengirim sinyal:
   cond.Signal()    → membangunkan satu goroutine
   cond.Broadcast() → membangunkan semua goroutine
```

---

## Catatan Penting

| Hal                              | Keterangan                                                                                           |
| -------------------------------- | ---------------------------------------------------------------------------------------------------- |
| **Lock sebelum Wait**            | `cond.L.Lock()` wajib dipanggil sebelum `cond.Wait()` — jika tidak, akan panic                       |
| **Wait melepas lock sementara**  | Selama menunggu, lock dilepas agar goroutine lain bisa masuk — lock diambil kembali saat dibangunkan |
| **`group.Add` sebelum `go`**     | Pastikan `Add(1)` dipanggil sebelum goroutine dijalankan, bukan di dalamnya                          |
| **Signal tidak menjamin urutan** | Goroutine mana yang dibangunkan oleh `Signal()` tidak bisa diprediksi                                |
| **Broadcast untuk semua**        | Gunakan `Broadcast()` jika semua goroutine yang menunggu perlu dibangunkan sekaligus                 |

---

Next: [Atomic](./23-atomic.md)
