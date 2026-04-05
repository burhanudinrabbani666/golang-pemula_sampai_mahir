# Time Ticker

## Gambaran Umum

**Ticker** adalah representasi kejadian yang **berulang** secara periodik. Berbeda dengan Timer yang hanya mengirim event sekali, Ticker terus mengirim event ke channel setiap kali durasinya habis — sampai eksplisit dihentikan dengan `Stop()`.

---

## Timer vs Ticker

| Aspek               | `time.Timer`      | `time.Ticker`                       |
| ------------------- | ----------------- | ----------------------------------- |
| Kejadian            | Satu kali         | Berulang terus                      |
| Dihentikan otomatis | ✅ Setelah expire | ❌ Harus dipanggil `Stop()`         |
| Cocok untuk         | Delay, timeout    | Polling, cron job, refresh periodik |

---

## Dua Cara Menggunakan Ticker

| Cara                | Fungsi                              | Cocok Untuk                                            |
| ------------------- | ----------------------------------- | ------------------------------------------------------ |
| `time.NewTicker(d)` | Membuat Ticker dengan kontrol penuh | Perlu bisa menghentikan ticker (`Stop()`)              |
| `time.Tick(d)`      | Langsung mengembalikan channel      | Ticker yang berjalan selamanya, tidak perlu dihentikan |

---

## 1. `time.NewTicker()`

Membuat Ticker dan mengekspos struct lengkapnya. Gunakan `ticker.Stop()` untuk menghentikannya — jika tidak, channel akan terus menerima data selamanya dan menyebabkan goroutine menggantung.

```go
func TestTicker(t *testing.T) {
    ticker := time.NewTicker(1 * time.Second)

    // Goroutine ini bertanggung jawab menghentikan ticker setelah 5 detik
    go func() {
        time.Sleep(5 * time.Second)
        ticker.Stop() // Hentikan ticker — channel ticker.C akan ditutup
    }()

    // Range dari channel akan berhenti otomatis setelah ticker.Stop() dipanggil
    // dan channel ticker.C ditutup
    for tick := range ticker.C {
        fmt.Println(tick) // Dicetak setiap 1 detik selama 5 detik
    }
}
```

**Contoh output:**

```
2024-01-01 10:00:01
2024-01-01 10:00:02
2024-01-01 10:00:03
2024-01-01 10:00:04
2024-01-01 10:00:05
```

---

## 2. `time.Tick()`

Shortcut dari `time.NewTicker()` — langsung mengembalikan channel tanpa struct Ticker. Karena tidak ada referensi ke Ticker-nya, ticker **tidak bisa dihentikan**. Gunakan hanya jika memang ingin ticker berjalan selamanya.

```go
func TestTick(t *testing.T) {
    channel := time.Tick(1 * time.Second)

    // Gunakan counter atau kondisi lain untuk menghentikan loop
    // karena ticker ini tidak bisa di-Stop()
    counter := 0
    for tick := range channel {
        fmt.Println(tick)
        counter++
        if counter == 5 {
            break // Keluar dari loop tanpa menghentikan ticker-nya
        }
    }
}
```

---

## Cara Kerja

```
time.NewTicker(1 * time.Second):

  t=0s → Ticker dibuat
  t=1s → kirim event ke ticker.C → fmt.Println()
  t=2s → kirim event ke ticker.C → fmt.Println()
  t=3s → kirim event ke ticker.C → fmt.Println()
  t=4s → kirim event ke ticker.C → fmt.Println()
  t=5s → ticker.Stop() dipanggil → channel ditutup → for range berhenti
```

---

## Mencegah Deadlock

Pertanyaan di kode asli: **"Bagaimana supaya tidak deadlock?"**

Jawabnya: `for range` pada channel akan **berhenti otomatis** saat channel ditutup. `ticker.Stop()` menutup channel `ticker.C`, sehingga loop `for range` akan keluar dengan sendirinya — tidak terjadi deadlock.

```
ticker.Stop()
  → channel ticker.C ditutup
  → for tick := range ticker.C berhenti
  → program lanjut setelah loop
```

Untuk `time.Tick()` yang tidak bisa dihentikan, gunakan `break` atau `select` dengan kondisi tertentu untuk keluar dari loop.

---

## Alur Penggunaan Ticker

```
1. Buat ticker: ticker := time.NewTicker(durasi)
        ↓
2. Jalankan goroutine untuk menghentikan ticker setelah kondisi tertentu
   go func() { time.Sleep(n); ticker.Stop() }()
        ↓
3. Iterasi channel: for tick := range ticker.C { ... }
        ↓
4. Loop berhenti otomatis saat ticker.Stop() dipanggil
```

---

## Catatan Penting

| Hal                                     | Keterangan                                                                       |
| --------------------------------------- | -------------------------------------------------------------------------------- |
| **Selalu panggil `Stop()`**             | Ticker yang tidak dihentikan akan bocor (goroutine dan memori tidak dibebaskan)  |
| **`time.Tick()` tidak bisa dihentikan** | Gunakan hanya untuk ticker yang memang seumur hidup program                      |
| **`for range` berhenti otomatis**       | Saat channel ditutup oleh `Stop()`, loop `for range` keluar — tidak ada deadlock |
| **Event mungkin tertunda**              | Jika goroutine sibuk, event dari ticker bisa antri di channel                    |
| **Bukan pengganti cron job**            | Ticker tidak menjamin waktu eksekusi yang presisi jika goroutine sedang dibebani |

---

Next: [GOMAXPROCS](./26-gomaxprocs.md)
