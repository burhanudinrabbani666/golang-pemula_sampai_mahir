# Sync WaitGroup

## Gambaran Umum

**`sync.WaitGroup`** digunakan untuk menunggu sekumpulan goroutine selesai sebelum program melanjutkan eksekusi. Ini adalah solusi yang lebih baik dibanding `time.Sleep()` untuk menunggu goroutine — karena `Sleep` hanya menebak berapa lama, sedangkan `WaitGroup` menunggu dengan pasti.

---

## Method WaitGroup

| Method         | Fungsi                                                                  |
| -------------- | ----------------------------------------------------------------------- |
| `group.Add(n)` | Menambahkan `n` ke counter — dipanggil **sebelum** goroutine dijalankan |
| `group.Done()` | Mengurangi counter sebesar 1 — dipanggil ketika goroutine **selesai**   |
| `group.Wait()` | Memblokir eksekusi sampai counter mencapai **0**                        |

> `Done()` setara dengan `Add(-1)`.

---

## Contoh Kode

```go
// RunAsynchronous menjalankan satu unit pekerjaan dalam goroutine.
// defer group.Done() memastikan counter selalu berkurang saat fungsi selesai,
// bahkan jika terjadi panic di tengah jalan.
func RunAsynchronous(group *sync.WaitGroup) {
    defer group.Done() // Kurangi counter saat fungsi ini selesai

    fmt.Println("Hello")
    time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
    group := &sync.WaitGroup{}

    for i := 0; i < 100; i++ {
        group.Add(1)              // Tambah counter SEBELUM goroutine dijalankan
        go RunAsynchronous(group) // Jalankan goroutine
    }

    group.Wait()          // Blokir di sini sampai semua 100 goroutine memanggil Done()
    fmt.Println("Selesai")
}
```

---

## Cara Kerja

```
group.Add(1) dipanggil 100 kali → counter = 100
        │
        ├── goroutine 1 selesai → Done() → counter = 99
        ├── goroutine 2 selesai → Done() → counter = 98
        ├── ...
        └── goroutine 100 selesai → Done() → counter = 0
                                                  │
                                           group.Wait() selesai blokir
                                                  │
                                          fmt.Println("Selesai") ✅
```

---

## Urutan yang Benar: `Add` Sebelum `go`

Ini adalah kesalahan umum yang menyebabkan race condition atau panic:

```go
// ❌ SALAH — Add dipanggil di dalam goroutine
go func() {
    group.Add(1) // Bisa terjadi setelah Wait() sudah dieksekusi
    defer group.Done()
    // ...
}()

// ✅ BENAR — Add dipanggil sebelum goroutine dijalankan
group.Add(1)
go func() {
    defer group.Done()
    // ...
}()
```

---

## Alur Penggunaan WaitGroup

```
1. Buat WaitGroup: group := &sync.WaitGroup{}
        ↓
2. Sebelum setiap goroutine: group.Add(1)
        ↓
3. Jalankan goroutine: go func() { ... }()
        ↓
4. Di dalam goroutine: defer group.Done()
        ↓
5. Setelah semua goroutine dijalankan: group.Wait()
        ↓
6. Lanjutkan kode setelah semua goroutine selesai
```

---

## Catatan Penting

| Hal                              | Keterangan                                                                  |
| -------------------------------- | --------------------------------------------------------------------------- |
| **`Add` sebelum `go`**           | Selalu panggil `Add()` sebelum memulai goroutine, bukan di dalamnya         |
| **`defer group.Done()`**         | Gunakan `defer` agar `Done()` selalu dipanggil meski goroutine panic        |
| **Counter tidak boleh negatif**  | Memanggil `Done()` lebih banyak dari `Add()` akan menyebabkan panic         |
| **Lebih baik dari `time.Sleep`** | `WaitGroup` menunggu dengan pasti, bukan menebak durasi                     |
| **Zero value siap pakai**        | `var group sync.WaitGroup` atau `&sync.WaitGroup{}` langsung bisa digunakan |

---

Next: [Sync Once](./19-sync-once.md)
