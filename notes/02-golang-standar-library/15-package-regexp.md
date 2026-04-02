# Package `regexp`

## Pengantar

Package `regexp` adalah utilitas bawaan Go untuk melakukan pencarian dan manipulasi **regular expression (regex)**. Go menggunakan library C buatan Google bernama **RE2** sebagai engine-nya, yang menjamin waktu eksekusi linear — artinya tidak ada risiko _catastrophic backtracking_ seperti pada beberapa engine regex lainnya.

**Referensi Resmi:**

- [Sintaks RE2](https://github.com/google/re2/wiki/Syntax)
- [Dokumentasi `regexp` — pkg.go.dev](https://golang.org/pkg/regexp/)

---

## Function Utama

| Function                                    | Return           | Kegunaan                                                             |
| ------------------------------------------- | ---------------- | -------------------------------------------------------------------- |
| `regexp.MustCompile(pattern string)`        | `*Regexp`        | Mengompilasi pattern regex. Panic jika pattern tidak valid           |
| `regexp.Compile(pattern string)`            | `*Regexp, error` | Seperti `MustCompile`, namun mengembalikan error alih-alih panic     |
| `Regexp.MatchString(s string)`              | `bool`           | Mengecek apakah string cocok dengan pattern                          |
| `Regexp.FindString(s string)`               | `string`         | Mengembalikan hasil match **pertama**                                |
| `Regexp.FindAllString(s string, n int)`     | `[]string`       | Mengembalikan semua hasil match (maks `n`; gunakan `-1` untuk semua) |
| `Regexp.FindStringSubmatch(s string)`       | `[]string`       | Mengembalikan match beserta **capture group**                        |
| `Regexp.ReplaceAllString(src, repl string)` | `string`         | Mengganti semua match dengan string pengganti                        |

> **Kapan pakai `MustCompile` vs `Compile`?**
> Gunakan `MustCompile` untuk pattern yang sudah pasti valid (hardcoded). Gunakan `Compile` jika pattern berasal dari input pengguna agar error bisa ditangani dengan aman.

---

## Contoh Penggunaan

### 1. MatchString — Cek Kecocokan

```go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    // Pattern: huruf 'e', diikuti satu huruf kecil, diikuti 'o'
    var regex *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)

    fmt.Println(regex.MatchString("eko")) // true  — cocok: e-k-o
    fmt.Println(regex.MatchString("edo")) // true  — cocok: e-d-o
    fmt.Println(regex.MatchString("eKo")) // false — 'K' bukan huruf kecil
}
```

### 2. FindAllString — Temukan Semua Match

```go
func main() {
    regex := regexp.MustCompile(`e([a-z])o`)

    input := "eko edo egi ego elo eto eKo"

    // Argumen kedua: batas maksimum hasil (-1 = semua)
    results := regex.FindAllString(input, -1)

    fmt.Println(results) // [eko edo ego elo eto]
}
```

### 3. FindStringSubmatch — Ambil Capture Group

```go
func main() {
    regex := regexp.MustCompile(`e([a-z])o`)

    match := regex.FindStringSubmatch("eko")
    fmt.Println(match[0]) // "eko" — full match
    fmt.Println(match[1]) // "k"   — isi capture group pertama
}
```

### 4. ReplaceAllString — Ganti Match

```go
func main() {
    regex := regexp.MustCompile(`e([a-z])o`)

    result := regex.ReplaceAllString("eko edo eKo", "***")
    fmt.Println(result) // "*** *** eKo"
}
```

---

## Sintaks Pattern Dasar

| Pattern           | Artinya                              |
| ----------------- | ------------------------------------ |
| `.`               | Sembarang karakter (kecuali newline) |
| `[a-z]`           | Huruf kecil a sampai z               |
| `[A-Z]`           | Huruf besar A sampai Z               |
| `[0-9]` atau `\d` | Digit 0 sampai 9                     |
| `\w`              | Word character (`[a-zA-Z0-9_]`)      |
| `\s`              | Whitespace                           |
| `+`               | Satu atau lebih                      |
| `*`               | Nol atau lebih                       |
| `?`               | Nol atau satu (opsional)             |
| `(abc)`           | Capture group                        |
| `^`               | Awal string                          |
| `$`               | Akhir string                         |

---

## Catatan Penting

- **RE2 tidak mendukung lookahead/lookbehind** — berbeda dengan PCRE (digunakan PHP, Python, dll). Jika butuh fitur tersebut, perlu pendekatan manual di Go.
- Pattern dikompilasi sekali lalu **digunakan berulang kali**; hindari memanggil `MustCompile` di dalam loop.
- `*Regexp` aman digunakan secara **concurrent** (goroutine-safe).

---

Next: [Package encoding](./16-package-encoding.md)
