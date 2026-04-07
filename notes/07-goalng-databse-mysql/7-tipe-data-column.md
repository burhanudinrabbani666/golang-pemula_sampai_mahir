# Tipe data column

- Sebelumnya kita hanya membuat table dengan tipe data di kolom nya berupa VARCHAR
- Untuk VARCHAR di database, biasanya kita gunakan String di Golang
- Bagaimana dengan tipe data yang lain?
- Apa representasinya di Golang, misal tipe data timestamp, date dan lain-lain

## Mapping Tipe Data

| Tipe Data Database              | Tipe Data Golang |
| ------------------------------- | ---------------- |
| VARCHAR, CHAR                   | string           |
| INT, BIGINT                     | int32, int64     |
| FLOAT, DOUBLE                   | float32, float64 |
| BOOLEAN                         | bool             |
| DATE, DATETIME, TIME, TIMESTAMP | time.Time        |

## Error Tipe Data Date

- Secara default, Driver MySQL untuk Golang akan melakukan query tipe data DATE, DATETIME, TIMESTAMP menjadi []byte / []uint8. Dimana ini bisa dikonversi menjadi String, lalu di parsing menjadi time.Time
- Namun hal ini merepotkan jika dilakukan manual, kita bisa meminta Driver MySQL untuk Golang secara otomatis melakukan parsing dengan menambahkan parameter parseDate=true

## Nullable Type

- Golang database tidak mengerti dengan tipe data NULL di database
- Oleh karena itu, khusus untuk kolom yang bisa NULL di database, akan jadi masalah jika kita melakukan Scan secara bulat-bulat menggunakan tipe data representasinya di Golang

## Error Data Null

- Konversi secara otomatis NULL tidak didukung oleh Driver MySQL Golang
- Oleh karena itu, khusus tipe kolom yang bisa NULL, kita perlu menggunakan tipe data yang ada dalam package sql

Next: [SQL injection](./8-sql-injection.md)
