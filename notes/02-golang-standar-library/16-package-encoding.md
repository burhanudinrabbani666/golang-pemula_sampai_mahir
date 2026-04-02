# Package `encoding`

## Pengantar

Package `encoding` menyediakan interface dan utilitas standar untuk melakukan **encode** dan **decode** data di Go. Di atasnya, Go menyediakan berbagai sub-package untuk format dan algoritma yang berbeda-beda.

**Referensi Resmi:**

- [Dokumentasi `encoding` — pkg.go.dev](https://pkg.go.dev/encoding)

---

## Sub-Package Encoding Populer

| Sub-Package       | Kegunaan                                                   |
| ----------------- | ---------------------------------------------------------- |
| `encoding/base64` | Encode/decode data biner ke representasi teks Base64       |
| `encoding/csv`    | Baca dan tulis data dalam format CSV                       |
| `encoding/json`   | Serialisasi dan deserialisasi data ke/dari format JSON     |
| `encoding/xml`    | Serialisasi dan deserialisasi data ke/dari format XML      |
| `encoding/hex`    | Encode/decode data ke representasi heksadesimal            |
| `encoding/binary` | Konversi tipe data ke/dari representasi biner (byte order) |

---

## Base64

Base64 digunakan untuk mengubah data biner menjadi string teks yang aman untuk ditransmisikan (misalnya melalui HTTP header, email, atau URL).

Go menyediakan beberapa varian encoding:

| Varian                  | Digunakan Untuk                                      |
| ----------------------- | ---------------------------------------------------- |
| `base64.StdEncoding`    | Encoding standar (RFC 4648), menggunakan `+` dan `/` |
| `base64.URLEncoding`    | Aman untuk URL, menggunakan `-` dan `_`              |
| `base64.RawStdEncoding` | Standar tanpa padding `=`                            |
| `base64.RawURLEncoding` | URL-safe tanpa padding `=`                           |

### Contoh: Encode dan Decode Base64

```go
package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    value := "Burhanudin Rabbani"

    // Encode: string → Base64
    encoded := base64.StdEncoding.EncodeToString([]byte(value))
    fmt.Println("Encoded:", encoded)
    // Output: QnVyaGFudWRpbiBSYWJiYW5p

    // Decode: Base64 → string
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Error:", err.Error())
    } else {
        fmt.Println("Decoded:", string(decoded))
        // Output: Burhanudin Rabbani
    }
}
```

> **Catatan:** `EncodeToString` menerima `[]byte`, bukan `string` langsung. Gunakan konversi `[]byte(value)` untuk mengubahnya.

---

## CSV

Package `encoding/csv` digunakan untuk membaca dan menulis data dalam format **Comma-Separated Values**. Format ini umum digunakan untuk pertukaran data tabular (spreadsheet, database export, dll).

### Contoh: Membaca CSV

`csv.NewReader` menerima `io.Reader` apa pun — bisa file, string, maupun network stream.

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "strings"
)

func main() {
    csvString := "Burhanudin,Rabbani,Aurelius\n" +
        "Ryan,Hidayat,Ganja\n" +
        "Heri,Pratama,Protein"

    reader := csv.NewReader(strings.NewReader(csvString))

    for {
        record, err := reader.Read()

        if err == io.EOF {
            // Akhir dari data CSV, hentikan loop
            break
        }
        if err != nil {
            fmt.Println("Error:", err)
            break
        }

        fmt.Println(record)
        // Output per baris: [Burhanudin Rabbani Aurelius], dst.
    }
}
```

### Contoh: Menulis CSV

`csv.NewWriter` menerima `io.Writer` apa pun — contoh di bawah menulis langsung ke `os.Stdout`, tapi bisa juga ke file.

```go
package main

import (
    "encoding/csv"
    "os"
)

func main() {
    writer := csv.NewWriter(os.Stdout)

    _ = writer.Write([]string{"Burhanudin", "Rabbani", "Aurelius"})
    _ = writer.Write([]string{"Ryan", "Hidayat", "AI"})
    _ = writer.Write([]string{"Heri", "Pratama", "Protein"})

    // Flush wajib dipanggil untuk memastikan semua data tertulis
    writer.Flush()

    if err := writer.Error(); err != nil {
        // Cek apakah ada error saat Flush
        panic(err)
    }
}
```

> **Penting:** `csv.Writer` menggunakan buffer internal. Tanpa memanggil `Flush()`, data bisa tidak tertulis sama sekali. Selalu cek `writer.Error()` setelah `Flush()` untuk memastikan tidak ada kegagalan I/O.

---

## Catatan Penting

- **Base64 bukan enkripsi** — ia hanya mengubah representasi data, bukan menyembunyikannya. Jangan gunakan Base64 untuk keamanan.
- **CSV tidak punya standar tunggal** — delimiter bisa berbeda (koma, titik koma, tab). Gunakan `reader.Comma = ';'` untuk mengubah delimiter di Go.
- Semua sub-package encoding di Go bekerja dengan `io.Reader` dan `io.Writer`, sehingga mudah dikomposisikan — misalnya membaca CSV langsung dari HTTP response body atau file terkompresi.

---

Next: [Package slices](./17-package-slices.md)
