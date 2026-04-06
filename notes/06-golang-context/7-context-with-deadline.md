# Context with Deadline

## Gambaran Umum

**`context.WithDeadline`** membuat context yang otomatis mengirim sinyal cancel pada **waktu absolut** tertentu. Berbeda dengan `WithTimeout` yang menghitung dari sekarang, `WithDeadline` menerima titik waktu spesifik sebagai batas.

```go
context.WithDeadline(parent, time.Time)   // waktu absolut
context.WithTimeout(parent, duration)     // relatif dari sekarang — setara dengan:
                                           // WithDeadline(parent, time.Now().Add(duration))
```

---

## Timeout vs Deadline — Perbedaan Praktis

| Aspek                 | `WithTimeout`                            | `WithDeadline`                    |
| --------------------- | ---------------------------------------- | --------------------------------- |
| Cara menentukan batas | Durasi dari sekarang                     | Waktu absolut                     |
| Contoh                | `5 * time.Second`                        | `time.Now().Add(5 * time.Second)` |
| Cocok untuk           | "Maksimal 5 detik dari sekarang"         | "Harus selesai sebelum jam 15:00" |
| Di balik layar        | Keduanya menggunakan mekanisme yang sama |                                   |

---

## Contoh Kode

```go
// CreateCounter versi lambat — ada jeda 1 detik per iterasi
// untuk mensimulasikan proses I/O yang memakan waktu
func CreateCounter(ctx context.Context) chan int {
    destination := make(chan int)

    go func() {
        defer close(destination)
        counter := 1

        for {
            select {
            case <-ctx.Done():
                return // Sinyal deadline/cancel diterima — berhenti

            default:
                destination <- counter
                counter++
                time.Sleep(1 * time.Second) // Simulasi proses lambat
            }
        }
    }()

    return destination
}

func TestContextWithDeadline(t *testing.T) {
    parent := context.Background()

    // Deadline: tepat 5 detik dari sekarang
    // Setara dengan: context.WithTimeout(parent, 5*time.Second)
    ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
    defer cancel() // Tetap wajib untuk membebaskan resource

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

    destination := CreateCounter(ctx)

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3

    // Karena ada Sleep 1 detik per counter, hanya ~5 angka yang tercetak
    // sebelum deadline tercapai dan goroutine berhenti
    for n := range destination {
        fmt.Println("Counter:", n)
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2
}
```

**Contoh output:**

```
Total Goroutine: 2
Total Goroutine: 3
Counter: 1
Counter: 2
Counter: 3
Counter: 4
Counter: 5
Total Goroutine: 2
```

---

## Cara Kerja

```
t=0s  → Deadline ditetapkan: time.Now() + 5s = jam 10:00:05
t=0s  → CreateCounter(ctx) dijalankan
t=1s  → Counter: 1
t=2s  → Counter: 2
t=3s  → Counter: 3
t=4s  → Counter: 4
t=5s  → Counter: 5
t=5s  → Deadline tercapai → ctx.Done() ditutup otomatis
t=5s  → goroutine CreateCounter menerima sinyal → return
t=5s  → channel destination ditutup → for range berhenti
t=7s  → NumGoroutine() = 2 ✅
```

---

## Alur Penggunaan Context with Deadline

```
1. Tentukan waktu deadline: deadline := time.Now().Add(durasi)
        ↓
2. Buat context: ctx, cancel := context.WithDeadline(parent, deadline)
        ↓
3. Selalu tambahkan: defer cancel()
        ↓
4. Teruskan ctx ke goroutine yang perlu dibatasi
        ↓
5. Goroutine memeriksa ctx.Done() — otomatis terpicu saat deadline tercapai
```

---

## Catatan Penting

| Hal                                    | Keterangan                                                                                    |
| -------------------------------------- | --------------------------------------------------------------------------------------------- |
| **Waktu absolut**                      | `WithDeadline` menerima `time.Time`, bukan durasi — gunakan `time.Now().Add()` untuk konversi |
| **`defer cancel()` tetap wajib**       | Membebaskan resource jika operasi selesai sebelum deadline                                    |
| **`ctx.Deadline()`**                   | Bisa digunakan untuk mengecek kapan deadline akan habis — mengembalikan `(time.Time, bool)`   |
| **Goroutine harus cek `ctx.Done()`**   | Deadline hanya mengirim sinyal — goroutine tetap harus mau berhenti sendiri                   |
| **Sama dengan Timeout di balik layar** | `WithTimeout(d)` secara internal memanggil `WithDeadline(time.Now().Add(d))`                  |
