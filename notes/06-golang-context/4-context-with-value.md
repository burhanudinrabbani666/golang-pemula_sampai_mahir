# Context with Value

## Gambaran Umum

Context bisa menyimpan data berupa pasangan **key-value**. Saat kita menambahkan value ke context, Go membuat **child context baru** — context aslinya tidak berubah sama sekali. Ini membentuk struktur **pohon (tree)** dari context yang saling berhubungan.

Untuk menambahkan value ke context, gunakan:

```go
context.WithValue(parent, key, value)
```

---

## Contoh Kode

### Membuat Pohon Context

```go
func TestContextWithValue(t *testing.T) {
    contextA := context.Background() // Root context — tidak memiliki value

    // Child langsung dari A
    contextB := context.WithValue(contextA, "b", "B")
    contextC := context.WithValue(contextA, "c", "C")

    // Child dari B
    contextD := context.WithValue(contextB, "d", "D")
    contextE := context.WithValue(contextB, "e", "E")

    // Child dari C
    contextF := context.WithValue(contextC, "f", "F")

    // Child dari F
    contextG := context.WithValue(contextF, "g", "G")

    fmt.Println(contextA)
    fmt.Println(contextB)
    fmt.Println(contextC)
    fmt.Println(contextD)
    fmt.Println(contextE)
    fmt.Println(contextF)
    fmt.Println(contextG)
}
```

---

## Struktur Pohon Context

```
contextA (root)
├── contextB  {b: "B"}
│   ├── contextD  {d: "D"}
│   └── contextE  {e: "E"}
└── contextC  {c: "C"}
    └── contextF  {f: "F"}
        └── contextG  {g: "G"}
```

Setiap node hanya "tahu" tentang dirinya sendiri dan **leluhurnya (ancestor)** — tidak tahu tentang saudaranya atau keturunannya.

---

## Mengambil Value

Value diambil menggunakan `context.Value(key)`. Pencarian dimulai dari context saat ini, lalu naik ke parent, ke grandparent, dan seterusnya sampai ditemukan atau sampai root.

```go
fmt.Println(contextF.Value("f")) // "F"  ← ditemukan di contextF sendiri
fmt.Println(contextF.Value("c")) // "C"  ← tidak ada di F, naik ke parent (contextC) → ditemukan
fmt.Println(contextF.Value("b")) // nil  ← tidak ada di F → C → A, jalur B tidak terhubung
fmt.Println(contextA.Value("b")) // nil  ← root tidak bisa mengambil value dari child
```

---

## Cara Kerja Pencarian Value

```
contextF.Value("c"):

  contextF  → cari key "c" → tidak ada
      ↓ naik ke parent
  contextC  → cari key "c" → ✅ ditemukan! kembalikan "C"

contextF.Value("b"):

  contextF  → cari key "b" → tidak ada
      ↓ naik ke parent
  contextC  → cari key "b" → tidak ada
      ↓ naik ke parent
  contextA  → cari key "b" → tidak ada (root)
  → kembalikan nil ← contextB adalah cabang berbeda, tidak bisa diakses
```

---

## Alur Penggunaan Context with Value

```
1. Buat root context: context.Background()
        ↓
2. Tambahkan value: ctx := context.WithValue(parent, key, value)
        ↓
3. Teruskan ctx ke fungsi atau goroutine yang membutuhkan
        ↓
4. Ambil value di mana saja: ctx.Value(key)
        ↓
5. Cek nil — key yang tidak ditemukan di jalur context akan mengembalikan nil
```

---

## Catatan Penting

| Hal                                  | Keterangan                                                                                     |
| ------------------------------------ | ---------------------------------------------------------------------------------------------- |
| **Immutable**                        | `WithValue` tidak mengubah parent — selalu membuat child context baru                          |
| **Pencarian ke atas**                | `Value()` mencari ke parent, grandparent, dst. — tidak pernah ke child atau sibling            |
| **Kembalikan nil jika tidak ada**    | Selalu cek apakah hasilnya nil sebelum digunakan                                               |
| **Key sebaiknya bukan string biasa** | Gunakan tipe khusus untuk key agar tidak tabrakan antar package (contoh: `type ctxKey string`) |
| **Jangan simpan data besar**         | Context dirancang untuk data request-scoped kecil seperti user ID, token, atau trace ID        |

---

Next: [Context with Cancel](./5-context-with-cancel.md)
