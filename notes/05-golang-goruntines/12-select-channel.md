# Select Channel

## Gambaran Umum

**Select channel** digunakan ketika kita memiliki beberapa channel yang berjalan secara paralel dan ingin menerima data dari **channel mana pun yang paling cepat** mengirim data.

Cara kerja `select` mirip seperti `switch`, namun khusus untuk operasi channel:

- Jika **satu channel** siap, maka `case` channel tersebut yang dijalankan
- Jika **beberapa channel** siap bersamaan, Go akan memilih salah satu secara **acak (random)**

---

## Sintaks Select

```go
select {
case data := <-channel1:
    // dijalankan jika channel1 mengirim data lebih dulu
case data := <-channel2:
    // dijalankan jika channel2 mengirim data lebih dulu
}
```

---

## Contoh Kode

```go
// GiveMeResponse adalah goroutine yang mengirim data ke channel setelah jeda waktu acak
func GiveMeResponse(channel chan string) {
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
    channel <- "Response dari goroutine"
}

func TestSelectChannel(t *testing.T) {
    channel1 := make(chan string)
    channel2 := make(chan string)

    defer close(channel1) // Tutup channel saat test selesai
    defer close(channel2)

    // Jalankan dua goroutine secara paralel
    go GiveMeResponse(channel1)
    go GiveMeResponse(channel2)

    counter := 0
    for {
        // select menunggu channel mana pun yang pertama kali siap
        select {
        case data := <-channel1:
            fmt.Println("Data dari Channel 1:", data)
            counter++

        case data := <-channel2:
            fmt.Println("Data dari Channel 2:", data)
            counter++
        }

        // Hentikan loop setelah kedua channel mengirim datanya
        if counter == 2 {
            break
        }
    }
}
```

---

## Cara Kerja

```
goroutine 1 → channel1 ──┐
                          ├──► select { menunggu yang tercepat }
goroutine 2 → channel2 ──┘           │
                                     ▼
                          channel1 lebih cepat? → jalankan case channel1
                          channel2 lebih cepat? → jalankan case channel2
                          keduanya bersamaan?   → pilih secara RANDOM

Loop terus sampai counter == 2 (semua channel sudah diproses)
```

---

## Alur Penggunaan Select Channel

```
1. Buat dua atau lebih channel
        ↓
2. Jalankan goroutine untuk masing-masing channel
        ↓
3. Gunakan for loop + select untuk menerima data
        ↓
4. Di dalam select, buat case untuk setiap channel
        ↓
5. Gunakan counter atau kondisi lain untuk menghentikan loop
        ↓
6. Tutup semua channel dengan defer close()
```

---

## Catatan Penting

| Hal                       | Keterangan                                                                                          |
| ------------------------- | --------------------------------------------------------------------------------------------------- |
| **Non-blocking**          | `select` hanya memilih channel yang **sudah siap** mengirim data                                    |
| **Random jika bersamaan** | Jika dua channel siap di waktu yang sama, Go memilih secara acak — tidak bisa diprediksi            |
| **Tanpa `for` loop**      | `select` hanya memproses **satu** channel per eksekusi — butuh loop jika ingin memproses semua      |
| **`counter`**             | Digunakan sebagai kondisi berhenti — pastikan nilainya sama dengan jumlah channel                   |
| **Tanpa `default`**       | `select` akan **block** sampai salah satu channel siap. Tambahkan `default` jika ingin non-blocking |

---

Next: [Default Select](./13-default-select.md)
