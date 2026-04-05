# Sync Pool

## Gambaran Umum

**`sync.Pool`** adalah implementasi dari design pattern **Object Pool** — sebuah pola di mana kita menyimpan sekumpulan objek yang siap dipakai, meminjamnya saat dibutuhkan, dan mengembalikannya setelah selesai — alih-alih membuat objek baru setiap kali diperlukan.

Ini berguna untuk menghemat biaya alokasi memori pada objek yang **mahal dibuat** namun **sering digunakan**, seperti koneksi database, buffer I/O, atau parser.

> `sync.Pool` sudah thread-safe — aman digunakan dari banyak goroutine secara bersamaan tanpa Mutex tambahan.

---

## Method Pool

| Method        | Fungsi                                                                        |
| ------------- | ----------------------------------------------------------------------------- |
| `pool.Put(x)` | Menyimpan objek `x` kembali ke pool untuk digunakan lagi                      |
| `pool.Get()`  | Mengambil objek dari pool — jika pool kosong, fungsi `New` dipanggil otomatis |

---

## Contoh Kode

```go
func TestPool(t *testing.T) {
    pool := sync.Pool{
        // New dipanggil otomatis oleh Get() jika pool sedang kosong
        New: func() any {
            return "New" // nilai default jika tidak ada objek tersedia di pool
        },
    }

    // Isi pool dengan 3 objek awal
    pool.Put("Burhanudin")
    pool.Put("D")
    pool.Put("Rabbani")

    // 10 goroutine berjalan — 3 pertama dapat objek dari pool,
    // sisanya akan mendapat nilai dari fungsi New karena pool kosong
    for i := 0; i < 10; i++ {
        go func() {
            data := pool.Get()         // Ambil objek dari pool
            fmt.Println(data)

            time.Sleep(1 * time.Second) // Simulasi penggunaan objek
            pool.Put(data)              // Kembalikan objek ke pool setelah selesai
        }()
    }

    time.Sleep(11 * time.Second)
    fmt.Println("Selesai")
}
```

**Contoh output (urutan bisa berbeda):**

```
Burhanudin
Rabbani
D
New
New
New
New
New
New
New
Selesai
```

---

## Cara Kerja

```
Pool awal: [ "Burhanudin", "D", "Rabbani" ]

Goroutine 1: Get() → dapat "Burhanudin" → pool: [ "D", "Rabbani" ]
Goroutine 2: Get() → dapat "D"          → pool: [ "Rabbani" ]
Goroutine 3: Get() → dapat "Rabbani"    → pool: [ ]
Goroutine 4: Get() → pool kosong → panggil New() → dapat "New"
Goroutine 5: Get() → pool kosong → panggil New() → dapat "New"
...

Setelah 1 detik, goroutine mulai Put() kembali:
Goroutine 1: Put("Burhanudin") → pool: [ "Burhanudin" ]
Goroutine 2: Put("D")          → pool: [ "Burhanudin", "D" ]
...
```

---

## Alur Penggunaan Sync Pool

```
1. Buat pool dengan fungsi New sebagai fallback
        ↓
2. (Opsional) Isi pool awal dengan Put()
        ↓
3. Saat butuh objek → pool.Get()
        ↓
4. Gunakan objek untuk keperluan tertentu
        ↓
5. Setelah selesai → pool.Put(objek) untuk dikembalikan
```

---

## Catatan Penting

| Hal                           | Keterangan                                                                                                     |
| ----------------------------- | -------------------------------------------------------------------------------------------------------------- |
| **`New` sebagai fallback**    | Dipanggil otomatis oleh `Get()` jika pool kosong — wajib diisi agar tidak panic                                |
| **Thread-safe**               | Aman digunakan dari banyak goroutine tanpa Mutex tambahan                                                      |
| **GC bisa mengosongkan pool** | Go's garbage collector bisa membersihkan isi pool kapan saja — jangan andalkan pool untuk penyimpanan permanen |
| **Selalu kembalikan objek**   | Jika `Put()` tidak dipanggil setelah `Get()`, pool akan terus memanggil `New()` dan membebani memori           |
| **Cocok untuk**               | Buffer, koneksi, parser, atau objek besar yang mahal untuk dibuat ulang                                        |

---

Next: [Sync Map](./21-sync-map.md)
