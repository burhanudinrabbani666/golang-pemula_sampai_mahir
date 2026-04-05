# Sync RWMutex

## Gambaran Umum

**`sync.RWMutex`** (Read Write Mutex) adalah pengembangan dari `sync.Mutex` yang memisahkan kunci untuk operasi **baca** dan **tulis**.

Dengan `Mutex` biasa, operasi baca dan tulis saling rebutan kunci — padahal membaca data secara bersamaan sebenarnya **aman** dilakukan oleh banyak goroutine sekaligus. `RWMutex` mengoptimalkan hal ini.

---

## Perbedaan Mutex vs RWMutex

| Kondisi                                | `sync.Mutex`                 | `sync.RWMutex`                                       |
| -------------------------------------- | ---------------------------- | ---------------------------------------------------- |
| Banyak goroutine **membaca** bersamaan | ❌ Harus antre satu per satu | ✅ Semua boleh baca bersamaan                        |
| Goroutine **menulis**                  | ✅ Dikunci eksklusif         | ✅ Dikunci eksklusif, semua baca/tulis lain menunggu |
| Cocok untuk                            | Operasi baca-tulis seimbang  | Operasi **baca jauh lebih sering** dari tulis        |

---

## Method RWMutex

| Method            | Fungsi                                                                              |
| ----------------- | ----------------------------------------------------------------------------------- |
| `mutex.Lock()`    | Kunci untuk **menulis** — semua goroutine lain (baca maupun tulis) harus menunggu   |
| `mutex.Unlock()`  | Lepas kunci tulis                                                                   |
| `mutex.RLock()`   | Kunci untuk **membaca** — goroutine baca lain **boleh** ikut masuk secara bersamaan |
| `mutex.RUnlock()` | Lepas kunci baca                                                                    |

---

## Contoh Kode: Rekening Bank

```go
package golanggoroutine

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

// BankAccount menggunakan RWMutex karena operasi baca (GetBalance)
// jauh lebih sering terjadi dibanding operasi tulis (AddBalance)
type BankAccount struct {
    RwMutex sync.RWMutex
    Balance int
}

// AddBalance mengubah saldo — gunakan Lock() karena operasi tulis (eksklusif)
func (account *BankAccount) AddBalance(amount int) {
    account.RwMutex.Lock()
    account.Balance = account.Balance + amount
    account.RwMutex.Unlock()
}

// GetBalance membaca saldo — gunakan RLock() agar banyak goroutine bisa baca bersamaan
func (account *BankAccount) GetBalance() int {
    account.RwMutex.RLock()
    balance := account.Balance
    account.RwMutex.RUnlock()
    return balance
}

func TestRWMutex(t *testing.T) {
    account := BankAccount{}

    // 100 goroutine berjalan paralel, masing-masing menambah saldo 100 kali
    for i := 0; i < 100; i++ {
        go func() {
            for j := 0; j < 100; j++ {
                account.AddBalance(1)              // Tulis: eksklusif
                fmt.Println(account.GetBalance())  // Baca: bisa paralel
            }
        }()
    }

    time.Sleep(5 * time.Second)
    fmt.Println("Total Balance:", account.GetBalance())
    // Hasil selalu konsisten: 10000 (100 goroutine × 100 iterasi × 1)
}
```

---

## Cara Kerja

```
Skenario: 3 goroutine membaca + 1 goroutine menulis

Goroutine A (baca):  RLock ✅ → baca balance → RUnlock
Goroutine B (baca):  RLock ✅ → baca balance → RUnlock  (boleh bersamaan dengan A)
Goroutine C (baca):  RLock ✅ → baca balance → RUnlock  (boleh bersamaan dengan A & B)
Goroutine D (tulis): Lock  ⏳ → tunggu A, B, C selesai → tulis balance → Unlock

Setelah D Unlock:
Goroutine berikutnya boleh RLock atau Lock lagi
```

---

## Alur Penggunaan RWMutex

```
1. Deklarasikan: var rwMutex sync.RWMutex (atau embed di struct)
        ↓
2. Untuk operasi TULIS:
   rwMutex.Lock()
   → ubah data
   rwMutex.Unlock()
        ↓
3. Untuk operasi BACA:
   rwMutex.RLock()
   → baca data
   rwMutex.RUnlock()
```

---

## Catatan Penting

| Hal                              | Keterangan                                                                             |
| -------------------------------- | -------------------------------------------------------------------------------------- |
| **Jangan campur Lock dan RLock** | Gunakan `Lock` hanya untuk tulis, `RLock` hanya untuk baca                             |
| **Writer menunggu semua Reader** | Goroutine tulis baru bisa masuk setelah **semua** goroutine baca selesai               |
| **Reader menunggu Writer**       | Jika ada goroutine tulis aktif, goroutine baca baru harus menunggu                     |
| **Jangan copy RWMutex**          | Selalu gunakan pointer atau embed langsung di struct                                   |
| **`defer` untuk Unlock**         | Gunakan `defer rwMutex.Unlock()` / `defer rwMutex.RUnlock()` agar kunci selalu dilepas |

---

Next: [Deadlock](./17-deadlock.md)
