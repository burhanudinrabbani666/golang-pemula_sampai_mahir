# Membuat Channel

## Gambaran Umum

**Channel** adalah mekanisme komunikasi antar goroutine di Go. Channel direpresentasikan dengan tipe data `chan` dan dibuat menggunakan fungsi `make()` — mirip seperti membuat map.

Saat membuat channel, kita harus menentukan **tipe data** yang akan dikirim melalui channel tersebut.

```go
channel := make(chan string) // channel khusus untuk tipe string
channel := make(chan int)    // channel khusus untuk tipe int
```

---

## Mengirim dan Menerima Data

| Operasi           | Sintaks             | Keterangan                                        |
| ----------------- | ------------------- | ------------------------------------------------- |
| **Kirim data**    | `channel <- data`   | Mengirim `data` ke dalam channel                  |
| **Terima data**   | `data := <-channel` | Mengambil data dari channel, disimpan ke variabel |
| **Tutup channel** | `close(channel)`    | Menutup channel setelah selesai digunakan         |

> ⚠️ **Penting:** Channel bersifat **blocking** — goroutine yang mengirim akan berhenti (block) sampai ada penerima, dan sebaliknya. Inilah yang membuat channel aman untuk sinkronisasi antar goroutine.

---

## Contoh Kode

```go
func TestCreateChannel(t *testing.T) {
    channel := make(chan string)
    defer close(channel) // Tutup channel saat fungsi selesai

    // Goroutine berjalan secara paralel di background
    go func() {
        time.Sleep(2 * time.Second)      // Simulasi proses yang memakan waktu
        channel <- "Burhanudin Rabbani"  // Kirim data — goroutine ini akan BLOCK di sini
                                         // sampai data diambil oleh penerima
        fmt.Println("Selesai mengirim data ke channel")
    }()

    // Main goroutine menunggu data dari channel (BLOCK sampai data tersedia)
    data := <-channel
    fmt.Println(data) // Output: Burhanudin Rabbani

    time.Sleep(5 * time.Second)
}
```

---

## Cara Kerja

```
Main goroutine                    Goroutine (background)
      |                                   |
      |                          Sleep 2 detik...
      |                                   |
      |                          channel <- "Burhanudin Rabbani"
      |                          [ BLOCK — menunggu penerima ]
      |                                   |
data := <-channel ←────────── data dikirim
      |                                   |
      |                          "Selesai mengirim data ke channel"
      |
fmt.Println(data)
→ Output: "Burhanudin Rabbani"
```

---

## Alur Penggunaan Channel

```
1. Buat channel dengan make(chan TipeData)
        ↓
2. Jalankan goroutine yang akan mengirim data (go func(){...}())
        ↓
3. Goroutine mengirim data: channel <- data
   [ goroutine BLOCK sampai ada penerima ]
        ↓
4. Main goroutine menerima data: data := <-channel
   [ main BLOCK sampai ada pengirim ]
        ↓
5. Tutup channel dengan close() setelah selesai
```

---

## Catatan Penting

| Hal                            | Keterangan                                                                                            |
| ------------------------------ | ----------------------------------------------------------------------------------------------------- |
| **Blocking by default**        | Pengirim dan penerima saling menunggu satu sama lain                                                  |
| **`defer close(channel)`**     | Selalu tutup channel setelah selesai untuk menghindari goroutine yang menggantung                     |
| **Tipe data wajib ditentukan** | `make(chan string)` hanya bisa diisi data bertipe `string`                                            |
| **Goroutine wajib**            | Pengiriman dan penerimaan harus dilakukan di goroutine yang berbeda, jika tidak akan terjadi deadlock |
| **Deadlock**                   | Jika pengirim dan penerima ada di goroutine yang sama dan saling menunggu, program akan crash         |

---

Next: [Channel sebagai Parameter](./8-channel-sebagai-parameter.md)
