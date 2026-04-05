# Time Timer

## Gambaran Umum

**Timer** merepresentasikan satu kejadian yang terjadi di masa depan. Ketika durasi timer habis (expire), sebuah event dikirim ke dalam **channel** milik timer tersebut. Go menyediakan tiga cara untuk menggunakan timer sesuai kebutuhan.

---

## Tiga Cara Menggunakan Timer

| Cara                   | Fungsi                                            | Cocok Untuk                                       |
| ---------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `time.NewTimer(d)`     | Membuat Timer dengan akses penuh ke struct-nya    | Perlu kontrol timer (bisa di-stop/reset)          |
| `time.After(d)`        | Langsung mengembalikan channel tanpa struct Timer | Hanya butuh channel-nya saja                      |
| `time.AfterFunc(d, f)` | Menjalankan fungsi setelah durasi tertentu        | Ingin jalankan fungsi dengan delay, tanpa channel |

---

## 1. `time.NewTimer()`

Membuat Timer dan mengekspos struct lengkapnya. Ketika timer expire, waktu saat itu dikirim ke `timer.C` (channel bertipe `<-chan time.Time`).

```go
func TestTimer(t *testing.T) {
    timer := time.NewTimer(5 * time.Second)

    fmt.Println(time.Now()) // Waktu saat ini (sebelum timer)

    timeEvent := <-timer.C  // Blokir sampai 5 detik berlalu
    fmt.Println(timeEvent)  // Waktu saat timer expire (sekitar 5 detik kemudian)
}
```

**Contoh output:**

```
2024-01-01 10:00:00
2024-01-01 10:00:05   ← 5 detik kemudian
```

---

## 2. `time.After()`

Shortcut dari `time.NewTimer()` — langsung mengembalikan channel tanpa perlu menyimpan struct Timer. Gunakan ini jika tidak perlu menghentikan atau mereset timer.

```go
func TestTimerAfter(t *testing.T) {
    channel := time.After(5 * time.Second) // Langsung dapat channel-nya

    fmt.Println(time.Now())

    timeEvent := <-channel  // Blokir sampai 5 detik berlalu
    fmt.Println(timeEvent)
}
```

---

## 3. `time.AfterFunc()`

Menjalankan sebuah fungsi di goroutine baru setelah durasi tertentu — tanpa perlu menangani channel sama sekali. Cocok ketika ingin menjalankan logika dengan delay.

```go
func TestAfterFunc(t *testing.T) {
    group := sync.WaitGroup{}
    group.Add(1)

    // Fungsi ini akan dijalankan otomatis setelah 5 detik di goroutine baru
    time.AfterFunc(5*time.Second, func() {
        fmt.Println("Selesai:", time.Now())
        group.Done()
    })

    fmt.Println("Mulai:", time.Now()) // Langsung dijalankan, tidak menunggu

    group.Wait() // Tunggu sampai fungsi AfterFunc selesai
}
```

**Contoh output:**

```
Mulai: 2024-01-01 10:00:00
Selesai: 2024-01-01 10:00:05   ← 5 detik kemudian
```

---

## Cara Kerja

```
time.NewTimer(5 * time.Second):

  t=0s  → Timer dibuat, mulai hitung mundur
  t=0s  → fmt.Println(time.Now()) dieksekusi
  t=0s  → <-timer.C dieksekusi → BLOKIR, menunggu...
  t=5s  → Timer expire → kirim waktu ke timer.C
  t=5s  → <-timer.C menerima data → lanjut eksekusi
  t=5s  → fmt.Println(timeEvent)

time.AfterFunc(5 * time.Second, func):

  t=0s  → Timer dibuat (non-blocking, langsung lanjut)
  t=0s  → fmt.Println("Mulai") dieksekusi
  t=5s  → Timer expire → func() dijalankan di goroutine baru
```

---

## Alur Penggunaan Timer

```
Butuh kontrol penuh (stop/reset)?
  → Gunakan time.NewTimer(d) → akses timer.C untuk menunggu

Hanya butuh channel-nya?
  → Gunakan time.After(d) → langsung tunggu dari channel

Ingin jalankan fungsi dengan delay?
  → Gunakan time.AfterFunc(d, func) → tidak perlu channel
```

---

## Catatan Penting

| Hal                                             | Keterangan                                                                      |
| ----------------------------------------------- | ------------------------------------------------------------------------------- |
| **`timer.C` adalah channel**                    | Bertipe `<-chan time.Time` — menerima waktu tepat saat timer expire             |
| **`time.After` tidak bisa di-stop**             | Tidak ada referensi ke Timer-nya, sehingga tidak bisa dihentikan sebelum expire |
| **`time.AfterFunc` berjalan di goroutine baru** | Fungsi yang dikirim tidak berjalan di goroutine yang sama dengan pemanggil      |
| **Gunakan WaitGroup dengan AfterFunc**          | Karena fungsinya async, perlu `WaitGroup` agar program tidak selesai lebih dulu |
| **`timer.Stop()`**                              | Hanya tersedia di `time.NewTimer()` — membatalkan timer sebelum expire          |

---

Next: [Time Ticker](./25-time-ticker.md)
