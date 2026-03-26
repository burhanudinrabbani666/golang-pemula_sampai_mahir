# Operasi boolean

| Operator | Keterangan |
| -------- | ---------- |
| &&       | Dan        |
|          | Atau       |
| !        | Kebalikan  |

## Operasi &&

| Nilai 1 | Operator | Nilai 2 | Hasil |
| ------- | -------- | ------- | ----- |
| true    | &&       | true    | true  |
| true    | &&       | false   | false |
| false   | &&       | true    | false |
| false   | &&       | false   | false |

## Operasi ||

OR == ||

| Nilai 1 | Operator | Nilai 2 | Hasil |
| ------- | -------- | ------- | ----- |
| true    | OR       | true    | true  |
| true    | OR       | false   | true  |
| false   | OR       | true    | true  |
| false   | OR       | false   | false |

## Operasi !

| Operator | Nilai 2 | Hasil |
| -------- | ------- | ----- |
| !        | true    | false |
| !        | false   | true  |

```go
func main() {
	nilaiAkhir := 90
	absensi := 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) // false
}
```

Next: [Tipe data array](./17-tipe-data-array.md)
