# Sync Mutex

## Gambaran Umum

**`sync.Mutex`** (Mutual Exclusion) adalah solusi bawaan Go untuk mengatasi race condition. Mutex bekerja seperti kunci pintu — hanya satu goroutine yang boleh masuk ke bagian kode kritis pada satu waktu. Goroutine lain harus menunggu sampai kunci dilepas.

---

## Cara Kerja Mutex

```
Goroutine 1: mutex.Lock() → masuk, kerjakan x = x + 1 → mutex.Unlock()
Goroutine 2: mutex.Lock() → ❌ TUNGGU (Goroutine 1 sedang pegang kunci)
                               ↓ (setelah Goroutine 1 Unlock)
Goroutine 2: mutex.Lock() → ✅ masuk, kerjakan x = x + 1 → mutex.Unlock()
```

Dengan mekanisme ini, tidak akan pernah ada dua goroutine yang memodifikasi `x` secara bersamaan.

---

## Method Mutex

| Method           | Fungsi                                                                                  |
| ---------------- | --------------------------------------------------------------------------------------- |
| `mutex.Lock()`   | Mengunci mutex — goroutine lain yang memanggil `Lock()` akan **block** sampai di-unlock |
| `mutex.Unlock()` | Melepas kunci — goroutine berikutnya yang menunggu diperbolehkan masuk                  |

---

## Contoh Kode

```go
func TestRaceConditionMutex(t *testing.T) {
    x := 0
    var mutex sync.Mutex // Zero value sudah siap digunakan, tidak perlu inisialisasi

    for i := 1; i < 1000; i++ {
        go func() {
            for j := 1; j <= 100; j++ {
                mutex.Lock()   // Kunci akses — goroutine lain harus menunggu di sini
                x = x + 1     // Bagian kritis: hanya satu goroutine yang bisa masuk
                mutex.Unlock() // Lepas kunci — goroutine berikutnya boleh masuk
            }
        }()
    }

    time.Sleep(5 * time.Second)
    fmt.Println("Counter =", x)
    // Dengan Mutex, hasilnya selalu konsisten: 99900
}
```

---

## Perbandingan: Tanpa vs Dengan Mutex

```
Tanpa Mutex:
  Goroutine A: baca x=5 → hitung 6 → tulis x=6
  Goroutine B: baca x=5 → hitung 6 → tulis x=6  ← lost update!
  Hasil: x=6 (seharusnya 7)

Dengan Mutex:
  Goroutine A: Lock → baca x=5 → hitung 6 → tulis x=6 → Unlock
  Goroutine B: Lock → (menunggu A selesai) → baca x=6 → hitung 7 → tulis x=7 → Unlock
  Hasil: x=7 ✅
```

---

## Alur Penggunaan Mutex

```
1. Deklarasikan: var mutex sync.Mutex
        ↓
2. Sebelum akses variabel shared → mutex.Lock()
        ↓
3. Lakukan operasi pada variabel (bagian kritis)
        ↓
4. Setelah selesai → mutex.Unlock()
        ↓
5. (Opsional) Gunakan defer mutex.Unlock() agar tidak lupa melepas kunci
```

---

## Catatan Penting

| Hal                        | Keterangan                                                                                  |
| -------------------------- | ------------------------------------------------------------------------------------------- |
| **Zero value siap pakai**  | `var mutex sync.Mutex` langsung bisa digunakan tanpa `make()` atau inisialisasi lain        |
| **Jangan copy Mutex**      | Selalu gunakan pointer (`*sync.Mutex`) jika diteruskan ke fungsi lain                       |
| **Deadlock**               | Jika `Lock()` dipanggil dua kali tanpa `Unlock()` di antaranya, program akan hang selamanya |
| **`defer mutex.Unlock()`** | Praktik terbaik agar kunci selalu dilepas meski terjadi error atau panic                    |
| **Tradeoff performa**      | Mutex memperlambat eksekusi karena goroutine harus antre — gunakan hanya di bagian kritis   |

---

Next: [Sync RWMutex](./16-sync-rwmutex.md)
