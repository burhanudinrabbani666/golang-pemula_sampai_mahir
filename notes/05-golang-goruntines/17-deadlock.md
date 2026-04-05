# Deadlock

## Gambaran Umum

**Deadlock** adalah kondisi di mana dua atau lebih goroutine saling menunggu kunci (lock) satu sama lain — sehingga tidak ada satu pun yang bisa melanjutkan eksekusi. Program tidak crash, tapi juga tidak pernah selesai.

> Deadlock adalah salah satu bug paling sulit ditemukan dalam pemrograman concurrent karena tidak selalu muncul setiap kali program dijalankan.

---

## Analogi Deadlock

Bayangkan dua orang di koridor sempit:

```
[ Bani ] →→→ menunggu Heri minggir
[ Heri ] →→→ menunggu Bani minggir

Keduanya saling menunggu → tidak ada yang bisa lewat → DEADLOCK
```

---

## Contoh Kode

### Struct dan Helper

```go
// UserBalance merepresentasikan saldo pengguna dengan proteksi Mutex
type UserBalance struct {
    Mutex   sync.Mutex
    Name    string
    Balance int
}

func (user *UserBalance) Lock()              { user.Mutex.Lock() }
func (user *UserBalance) Unlock()            { user.Mutex.Unlock() }
func (user *UserBalance) Change(amount int)  { user.Balance = user.Balance + amount }
```

### Fungsi Transfer — Sumber Deadlock

```go
// Transfer memindahkan sejumlah dana dari user1 ke user2.
// ⚠️ Fungsi ini berpotensi deadlock jika dipanggil dari dua goroutine
// dengan urutan user yang terbalik secara bersamaan.
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
    user1.Lock()
    fmt.Println("Lock user1:", user1.Name)
    user1.Change(-amount)

    time.Sleep(1 * time.Second) // Simulasi proses — memberi waktu goroutine lain masuk

    user2.Lock() // ⚠️ Goroutine bisa BLOCK di sini selamanya jika user2 sudah dikunci goroutine lain
    fmt.Println("Lock user2:", user2.Name)
    user2.Change(amount)

    time.Sleep(1 * time.Second)

    user1.Unlock()
    user2.Unlock()
}
```

### Test Deadlock

```go
func TestDeadlock(t *testing.T) {
    user1 := UserBalance{Name: "Bani", Balance: 1000000}
    user2 := UserBalance{Name: "Heri", Balance: 1000000}

    // Goroutine A: Transfer user1 → user2 (kunci user1 dulu, lalu user2)
    go Transfer(&user1, &user2, 100000)

    // Goroutine B: Transfer user2 → user1 (kunci user2 dulu, lalu user1)
    // Urutan kunci TERBALIK dari Goroutine A → inilah penyebab deadlock
    go Transfer(&user2, &user1, 200000)

    time.Sleep(10 * time.Second)

    fmt.Println("User:", user1.Name, "Balance:", user1.Balance)
    fmt.Println("User:", user2.Name, "Balance:", user2.Balance)
}
```

---

## Cara Kerja (Mengapa Terjadi Deadlock)

```
Goroutine A (Transfer user1 → user2)   Goroutine B (Transfer user2 → user1)
            │                                       │
  Lock user1 ✅                           Lock user2 ✅
  (pegang kunci user1)                    (pegang kunci user2)
            │                                       │
  Sleep 1 detik...                        Sleep 1 detik...
            │                                       │
  Lock user2 ⏳                           Lock user1 ⏳
  (user2 dikunci B → TUNGGU)              (user1 dikunci A → TUNGGU)
            │                                       │
            └───────────── DEADLOCK ────────────────┘
            Keduanya menunggu satu sama lain selamanya
```

---

## Cara Mencegah Deadlock

### 1. Konsistenkan Urutan Kunci

Selalu kunci resource dalam **urutan yang sama** di semua goroutine:

```go
// ✅ Selalu kunci dengan urutan yang konsisten (misal: berdasarkan nama atau ID)
func TransferSafe(from, to *UserBalance, amount int) {
    // Tentukan urutan kunci berdasarkan nama agar konsisten
    first, second := from, to
    if from.Name > to.Name {
        first, second = to, from
    }

    first.Lock()
    defer first.Unlock()
    second.Lock()
    defer second.Unlock()

    from.Change(-amount)
    to.Change(amount)
}
```

### 2. Gunakan `defer` untuk Unlock

```go
user1.Lock()
defer user1.Unlock() // Kunci pasti dilepas, bahkan jika terjadi panic
```

---

## Alur Penggunaan (Pola Aman)

```
1. Tentukan urutan penguncian resource yang konsisten di seluruh kode
        ↓
2. Selalu gunakan defer Unlock() setelah Lock()
        ↓
3. Hindari memanggil Lock() di dalam Lock() pada resource yang sama
        ↓
4. Batasi waktu tunggu dengan timeout jika memungkinkan
        ↓
5. Uji dengan go test -race untuk mendeteksi potensi masalah
```

---

## Catatan Penting

| Hal                      | Keterangan                                                                                                 |
| ------------------------ | ---------------------------------------------------------------------------------------------------------- |
| **Program tidak crash**  | Deadlock membuat program hang (diam), bukan panic — sulit dideteksi                                        |
| **Urutan kunci**         | Penyebab paling umum deadlock adalah urutan penguncian yang tidak konsisten                                |
| **`defer Unlock()`**     | Mencegah kunci tidak pernah dilepas akibat lupa atau panic                                                 |
| **Go detector**          | Go runtime mendeteksi deadlock sederhana dan mencetak `fatal error: all goroutines are asleep - deadlock!` |
| **Deadlock tersembunyi** | Deadlock yang melibatkan banyak goroutine bisa tidak terdeteksi otomatis oleh runtime                      |

---

Next: [Sync WaitGroup](./18-sync-waitgroup.md)
