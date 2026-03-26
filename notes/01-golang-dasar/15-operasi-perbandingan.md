# Operasi perbandingan

| Operator | Keterangan        |
| -------- | ----------------- |
| ++       | a = a + 1         |
| --       | a = a - 1         |
| -        | Negative          |
| +        | Positive          |
| !        | Boolean kebalikan |

```go
func main() {
	var j = 1
	j++
	j++

	fmt.Println(j) // 3
}
```

- Operasi perbandingan adalah operasi untuk membandingkan dua buah data
- Operasi perbandingan adalah operasi yang menghasilkan nilai boolean (benar atau salah)
- Jika hasil operasinya adalah benar, maka nilainya adalah true
- Jika hasil operasinya adalah salah, maka nilainya adalah false

| Operator | Keterangan              |
| -------- | ----------------------- |
| >        | Lebih Dari              |
| <        | Kurang Dari             |
| >=       | Lebih Dari Sama Dengan  |
| <=       | Kurang Dari Sama Dengan |
| ==       | Sama Dengan             |
| !=       | Tidak Sama Dengan       |

```go
func main() {
	var name1 = "bani"
	var name2 = "bani"

	var result bool = name1 == name2

	fmt.Println(result) // true

}
```

Next: [Operasi boolean](./16-operasi-boolean.md)
