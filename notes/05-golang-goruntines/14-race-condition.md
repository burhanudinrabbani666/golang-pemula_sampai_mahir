# Race Condition

## Gambaran Umum

**Race condition** adalah kondisi di mana dua atau lebih goroutine mengakses dan memodifikasi variabel yang sama secara bersamaan, menghasilkan nilai akhir yang tidak dapat diprediksi.

Go menjalankan goroutine tidak hanya secara **concurrent** (bergantian), tetapi bisa juga **parallel** (benar-benar bersamaan di beberapa CPU thread). Inilah yang membuat race condition bisa terjadi.

---

## Mengapa Ini Berbahaya?

Bayangkan dua goroutine membaca nilai `x = 5` secara bersamaan, lalu keduanya menghitung `x + 1 = 6`, dan masing-masing menulis `6` kembali ke `x`. Padahal seharusnya hasil akhirnya adalah `7` (dua kali penambahan).

```
Goroutine A: baca x=5 → hitung 5+1=6 → tulis x=6
Goroutine B: baca x=5 → hitung 5+1=6 → tulis x=6
                                               ↑
                               Seharusnya x=7, tapi hasilnya x=6 ← BUG
```

Ini yang disebut **lost update** — salah satu penambahan hilang karena kedua goroutine membaca nilai lama secara bersamaan.

---

## Contoh Kode

```go
func TestRaceCondition(t *testing.T) {
    x := 0

    // Jalankan 999 goroutine secara paralel
    for i := 1; i < 1000; i++ {
        go func() {
            // Setiap goroutine menambah x sebanyak 100 kali
            for j := 1; j <= 100; j++ {
                x = x + 1 // ⚠️ Tidak aman! Banyak goroutine mengakses x bersamaan
            }
        }()
    }

    time.Sleep(5 * time.Second) // Tunggu semua goroutine selesai (tidak disarankan untuk production)
    fmt.Println("Counter =", x)
    // Nilai x seharusnya 999 * 100 = 99.900
    // Namun karena race condition, hasilnya bisa berbeda setiap kali dijalankan
}
```

**Hasil yang diharapkan:** `99900`

**Hasil yang mungkin terjadi:** `87342`, `91205`, atau angka lain yang tidak konsisten — berbeda setiap kali program dijalankan.

---

## Cara Kerja (Mengapa Nilai Bisa Salah)

```
Nilai awal: x = 0

Goroutine 1          Goroutine 2          Goroutine 3
    │                    │                    │
baca x=0            baca x=0            baca x=0
hitung 0+1=1        hitung 0+1=1        hitung 0+1=1
tulis x=1           tulis x=1           tulis x=1
                                              │
                              x seharusnya = 3, tapi hasilnya = 1
                              → 2 penambahan hilang (lost update)
```

Semakin banyak goroutine berjalan paralel, semakin banyak update yang hilang.

---

## Mendeteksi Race Condition

Go menyediakan **race detector** bawaan yang bisa diaktifkan saat menjalankan test:

```bash
go test -race ./...
```

Jika ada race condition, Go akan mencetak peringatan lengkap beserta lokasi kode yang bermasalah.

---

## Solusi Race Condition

| Solusi            | Kapan Digunakan                                   |
| ----------------- | ------------------------------------------------- |
| **`sync.Mutex`**  | Melindungi akses ke variabel shared (paling umum) |
| **`sync/atomic`** | Operasi sederhana seperti increment counter       |
| **Channel**       | Komunikasi antar goroutine tanpa shared variable  |

---

## Catatan Penting

| Hal                           | Keterangan                                                                     |
| ----------------------------- | ------------------------------------------------------------------------------ |
| **Tidak deterministik**       | Hasil race condition berbeda setiap kali program dijalankan                    |
| **`time.Sleep` bukan solusi** | Menunggu dengan Sleep tidak menjamin semua goroutine selesai dengan benar      |
| **Race detector**             | Gunakan `go test -race` untuk mendeteksi race condition secara otomatis        |
| **Shared variable**           | Hindari memodifikasi variabel yang sama dari beberapa goroutine tanpa proteksi |

---

Next: [Sync Mutex](./15-sync-mutex.md)
