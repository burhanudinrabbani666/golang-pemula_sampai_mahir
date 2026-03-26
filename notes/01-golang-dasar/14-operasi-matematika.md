# Operasi matematika

| Operator | Keterangan     |
| -------- | -------------- |
| +        | Penjumlahan    |
| -        | Pengurangan    |
| \*       | Perkalian      |
| /        | Pembagian      |
| %        | Sisa Pembagian |

```go
func main() {
	a := 10
	b := 10
	c := a + b

	fmt.Println(c)
}
```

| Operasi Matematika | Augmented Assignments |
| ------------------ | --------------------- |
| a = a + 10         | a += 10               |
| a = a - 10         | a -= 10               |
| a = a \* 10        | a \*= 10              |
| a = a / 10         | a /= 10               |
| a = a % 10         | a %= 10               |

```go
func main() {

	var a = 10
	a += 10

	fmt.Println(a) // 20

	a += 5
	fmt.Println(a) // 25

}
```

| Operator | Keterangan        |
| -------- | ----------------- |
| ++       | a = a + 1         |
| --       | a = a - 1         |
| -        | Negative          |
| +        | Positive          |
| !        | Boolean kebalikan |

Next: [Operasi perbandingan](./15-operasi-perbandingan.md)
