# Context with Timeout

## Gambaran Umum

**`context.WithTimeout`** adalah cara membuat context yang otomatis mengirim **sinyal cancel setelah durasi tertentu** — tanpa perlu memanggil `cancel()` secara manual. Ini sangat cocok untuk operasi yang memiliki batas waktu maksimal, seperti query database atau pemanggilan HTTP API.

---

## Perbandingan: Cancel vs Timeout vs Deadline

| Fungsi                 | Kapan Cancel Dipicu                         | Cocok Untuk                                 |
| ---------------------- | ------------------------------------------- | ------------------------------------------- |
| `context.WithCancel`   | Manual — saat `cancel()` dipanggil          | Membatalkan berdasarkan kondisi logika      |
| `context.WithTimeout`  | Otomatis — setelah durasi dari **sekarang** | Batas waktu relatif (5 detik dari sekarang) |
| `context.WithDeadline` | Otomatis — pada waktu **absolut** tertentu  | Batas waktu spesifik (jam 15:00 tepat)      |

---

## Contoh Kode

```go
func TestContextWithTimeout(t *testing.T) {
    parent := context.Background()

    // Context ini akan otomatis mengirim sinyal cancel setelah 5 detik
    ctx, cancel := context.WithTimeout(parent, 5*time.Second)
    defer cancel() // Tetap perlu dipanggil untuk membebaskan resource jika selesai lebih awal

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

    destination := CreateCounter(ctx) // Goroutine CreateCounter menerima ctx

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3

    // Loop ini akan berhenti otomatis saat timeout 5 detik habis
    // karena ctx.Done() akan terpicu dan goroutine CreateCounter akan berhenti,
    // menutup channel destination
    for n := range destination {
        fmt.Println("Counter:", n)
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2 — goroutine sudah berhenti
}
```

> Fungsi `CreateCounter(ctx)` yang digunakan di sini adalah versi dengan `select` + `ctx.Done()` dari materi sebelumnya — bukan versi tanpa context.

---

## Cara Kerja

```
t=0s  → context.WithTimeout(parent, 5s) dibuat
t=0s  → CreateCounter(ctx) dijalankan → goroutine aktif, mulai kirim data
t=0s  → for range destination mulai menerima: 1, 2, 3, ...
t=5s  → timeout habis → ctx.Done() otomatis ditutup
t=5s  → goroutine CreateCounter menerima sinyal dari ctx.Done() → return
t=5s  → channel destination ditutup → for range berhenti
t=7s  → NumGoroutine() → 2 (goroutine sudah bersih)
```

---

## Mengapa `defer cancel()` Tetap Diperlukan?

Meskipun timeout sudah otomatis, `cancel()` tetap harus dipanggil:

```
Jika program selesai SEBELUM timeout habis:
  → Tanpa defer cancel(), resource timer di dalam context tidak langsung dibebaskan
  → Dengan defer cancel(), resource langsung dibersihkan saat fungsi selesai

Aturan: Selalu panggil defer cancel() setelah context.WithTimeout()
```

---

## Alur Penggunaan Context with Timeout

```
1. Buat context: ctx, cancel := context.WithTimeout(parent, durasi)
        ↓
2. Selalu tambahkan: defer cancel()
        ↓
3. Teruskan ctx ke goroutine atau fungsi yang perlu dibatasi waktunya
        ↓
4. Goroutine memeriksa ctx.Done() — otomatis terpicu saat timeout habis
        ↓
5. Tidak perlu memanggil cancel() manual — timeout menanganinya
```

---

## Catatan Penting

| Hal                                  | Keterangan                                                                               |
| ------------------------------------ | ---------------------------------------------------------------------------------------- |
| **Cancel otomatis**                  | Sinyal cancel dikirim otomatis setelah durasi habis — tidak perlu dipanggil manual       |
| **`defer cancel()` tetap wajib**     | Mencegah resource leak jika operasi selesai sebelum timeout                              |
| **Goroutine harus cek `ctx.Done()`** | Timeout hanya mengirim sinyal — goroutine tetap harus mau berhenti sendiri               |
| **Cocok untuk I/O**                  | Ideal untuk query database, HTTP request, atau operasi jaringan dengan batas waktu       |
| **Timeout relatif**                  | Dihitung dari waktu `WithTimeout` dipanggil — gunakan `WithDeadline` untuk waktu absolut |

---

Next: [Context with Deadline](./7-context-with-deadline.md)
