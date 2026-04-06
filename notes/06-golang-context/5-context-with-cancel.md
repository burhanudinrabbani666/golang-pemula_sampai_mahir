# Context with Cancel

## Gambaran Umum

**`context.WithCancel`** memungkinkan kita mengirim **sinyal pembatalan** ke goroutine lain yang sedang berjalan. Goroutine penerima bertanggung jawab untuk memeriksa sinyal tersebut dan menghentikan dirinya sendiri.

> Goroutine tidak bisa dihentikan paksa dari luar — ia harus mau berhenti sendiri dengan memeriksa `ctx.Done()`.

---

## Masalah: Goroutine Leak

Tanpa context, goroutine yang berjalan terus-menerus tidak bisa dihentikan dari luar — inilah yang disebut **goroutine leak**.

```go
// ❌ Versi bermasalah — goroutine tidak bisa dihentikan
func CreateCounter() chan int {
    destination := make(chan int)

    go func() {
        defer close(destination)
        counter := 1

        for {
            destination <- counter // Goroutine ini berjalan selamanya
            counter++
        }
    }()

    return destination
}

func TestContextWithCancel(t *testing.T) {
    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

    destination := CreateCounter()

    for n := range destination {
        fmt.Println("Counter:", n)
        if n == 10 {
            break // Keluar dari loop, tapi goroutine di dalam CreateCounter MASIH BERJALAN
        }
    }

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // Masih 2 — goroutine bocor!
}
```

Setelah `break`, goroutine di dalam `CreateCounter` masih hidup karena tidak ada yang memberi tahu ia harus berhenti.

---

## Solusi: Context with Cancel

```go
// ✅ Versi dengan context — goroutine bisa dihentikan dari luar
func CreateCounter(ctx context.Context) chan int {
    destination := make(chan int)

    go func() {
        defer close(destination)
        counter := 1

        for {
            select {
            case <-ctx.Done():
                // Sinyal cancel diterima — hentikan goroutine
                return

            default:
                // Belum ada sinyal cancel — lanjut kirim data
                destination <- counter
                counter++
            }
        }
    }()

    return destination
}

func TestContextWithCancel(t *testing.T) {
    parent := context.Background()
    ctx, cancel := context.WithCancel(parent) // cancel adalah fungsi untuk mengirim sinyal stop

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2

    destination := CreateCounter(ctx)

    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 3 (goroutine CreateCounter aktif)

    for n := range destination {
        fmt.Println("Counter:", n)
        if n == 10 {
            break
        }
    }

    cancel() // Kirim sinyal cancel → ctx.Done() channel akan ditutup

    time.Sleep(2 * time.Second)
    fmt.Println("Total Goroutine:", runtime.NumGoroutine()) // 2 (goroutine sudah berhenti)
}
```

---

## Cara Kerja

```
context.WithCancel(parent)
  → mengembalikan (ctx, cancel)
  → ctx.Done() adalah channel yang akan ditutup saat cancel() dipanggil

Goroutine berjalan:
  for {
    select {
      case <-ctx.Done() → cancel() sudah dipanggil → return (berhenti)
      default           → belum ada sinyal → lanjut kerja
    }
  }

Saat cancel() dipanggil:
  → ctx.Done() channel ditutup
  → case <-ctx.Done() di goroutine terpicu
  → goroutine return → selesai
```

---

## Perbandingan: Tanpa vs Dengan Cancel

```
Tanpa Cancel:
  Goroutine dibuat → berjalan selamanya → tidak bisa dihentikan → GOROUTINE LEAK

Dengan Cancel:
  Goroutine dibuat → cek ctx.Done() setiap iterasi
  → cancel() dipanggil → goroutine berhenti sendiri → tidak ada leak
```

---

## Alur Penggunaan Context with Cancel

```
1. Buat context: ctx, cancel := context.WithCancel(parent)
        ↓
2. Teruskan ctx ke goroutine yang perlu bisa dibatalkan
        ↓
3. Di dalam goroutine, cek ctx.Done() menggunakan select
        ↓
4. Saat ingin membatalkan: panggil cancel()
        ↓
5. Goroutine menerima sinyal dari ctx.Done() dan berhenti sendiri
```

---

## Catatan Penting

| Hal                                   | Keterangan                                                                              |
| ------------------------------------- | --------------------------------------------------------------------------------------- |
| **Goroutine harus cek sendiri**       | `cancel()` hanya mengirim sinyal — goroutine yang memutuskan kapan berhenti             |
| **Selalu panggil `cancel()`**         | Jika tidak dipanggil, resource context tidak akan dibebaskan (gunakan `defer cancel()`) |
| **`ctx.Done()` adalah channel**       | Ditutup saat `cancel()` dipanggil — bisa digunakan di `select`                          |
| **Cancel berlaku ke seluruh subtree** | Memanggil `cancel()` juga membatalkan semua child context yang dibuat dari ctx          |
| **Goroutine leak berbahaya**          | Goroutine yang tidak pernah berhenti terus memakan memori selama program berjalan       |

---

Next: [Context with Timeout](./6-context-with-timeout.md)
