# GOMAXPROCS

## Gambaran Umum

Goroutine berjalan di dalam **OS Thread**. Go secara otomatis mengelola berapa banyak thread yang aktif menggunakan nilai **GOMAXPROCS** — yaitu jumlah thread OS yang bisa menjalankan goroutine secara paralel pada satu waktu.

Secara default, nilai GOMAXPROCS sama dengan **jumlah CPU** di komputer yang menjalankan program.

---

## Tiga Fungsi Penting di Package `runtime`

| Fungsi                   | Nilai Kembalian          | Keterangan                                                            |
| ------------------------ | ------------------------ | --------------------------------------------------------------------- |
| `runtime.NumCPU()`       | `int`                    | Jumlah logical CPU yang tersedia                                      |
| `runtime.GOMAXPROCS(n)`  | `int` (nilai sebelumnya) | Set jumlah thread OS; gunakan `-1` untuk hanya membaca tanpa mengubah |
| `runtime.NumGoroutine()` | `int`                    | Jumlah goroutine yang sedang aktif saat ini                           |

---

## Contoh Kode

### Membaca Informasi Runtime

```go
func TestGetGomaxprocs(t *testing.T) {
    group := sync.WaitGroup{}

    // Jalankan 100 goroutine agar NumGoroutine() menunjukkan angka yang terlihat
    for i := 0; i < 100; i++ {
        group.Go(func() {
            time.Sleep(3 * time.Second)
        })
    }

    totalCPU := runtime.NumCPU()
    fmt.Println("Total CPU:", totalCPU) // Contoh output: 8

    // GOMAXPROCS(-1) hanya membaca nilai saat ini tanpa mengubahnya
    totalThread := runtime.GOMAXPROCS(-1)
    fmt.Println("Total Thread:", totalThread) // Sama dengan jumlah CPU secara default

    totalGoRoutine := runtime.NumGoroutine()
    fmt.Println("Total Goroutine:", totalGoRoutine) // 100 goroutine + 1 goroutine utama = 101

    group.Wait()
}
```

**Contoh output (pada mesin dengan 8 CPU):**

```
Total CPU: 8
Total Thread: 8
Total Goroutine: 101
```

---

### Mengubah Jumlah Thread

```go
func TestChangeThreadNum(t *testing.T) {
    group := sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        group.Go(func() {
            time.Sleep(3 * time.Second)
        })
    }

    totalCPU := runtime.NumCPU()
    fmt.Println("Total CPU:", totalCPU)

    // Mengubah jumlah thread OS menjadi 20
    // ⚠️ Jarang diperlukan — Go sudah mengoptimalkan nilai default secara otomatis
    runtime.GOMAXPROCS(20)

    totalThread := runtime.GOMAXPROCS(-1) // Baca nilai setelah diubah
    fmt.Println("Total Thread:", totalThread) // Output: 20

    totalGoRoutine := runtime.NumGoroutine()
    fmt.Println("Total Goroutine:", totalGoRoutine)

    group.Wait()
}
```

---

## Cara Kerja

```
Komputer dengan 8 CPU:

  GOMAXPROCS = 8 (default)
  → Maksimal 8 goroutine berjalan BENAR-BENAR paralel pada satu waktu
  → Goroutine lainnya menunggu giliran di antrian scheduler Go

  100 goroutine dibuat:
  ┌─────────────────────────────────────────┐
  │ Thread 1 │ Thread 2 │ ... │ Thread 8   │  ← berjalan paralel
  │ Goroutine│ Goroutine│     │ Goroutine  │
  └─────────────────────────────────────────┘
  Sisa 92 goroutine menunggu di antrian scheduler
```

---

## Kapan Perlu Mengubah GOMAXPROCS?

```
Umumnya: TIDAK PERLU — Go scheduler sudah sangat optimal

Kasus khusus yang mungkin perlu diubah:
  - Membatasi penggunaan CPU di lingkungan shared/container
  - Eksperimen performa untuk profiling
  - Aplikasi dengan karakteristik I/O bound vs CPU bound yang ekstrem
```

---

## Alur Penggunaan

```
1. Baca jumlah CPU: runtime.NumCPU()
        ↓
2. Baca GOMAXPROCS saat ini: runtime.GOMAXPROCS(-1)
        ↓
3. (Opsional) Ubah jika diperlukan: runtime.GOMAXPROCS(n)
        ↓
4. Monitor goroutine aktif: runtime.NumGoroutine()
```

---

## Catatan Penting

| Hal                              | Keterangan                                                                              |
| -------------------------------- | --------------------------------------------------------------------------------------- |
| **Default = jumlah CPU**         | Go otomatis menyesuaikan GOMAXPROCS dengan jumlah logical CPU                           |
| **GOMAXPROCS(-1)**               | Membaca nilai saat ini tanpa mengubahnya — nilai kembalian adalah angka thread aktif    |
| **Goroutine ≠ Thread**           | Ribuan goroutine bisa berjalan di atas hanya beberapa thread OS                         |
| **Jarang perlu diubah**          | Go scheduler modern sudah sangat efisien — ubah hanya jika ada alasan performa spesifik |
| **`NumGoroutine` termasuk main** | Nilai mencakup goroutine utama program, bukan hanya yang dibuat manual                  |
