# Table benchmark

- Sama seperti di unit test, programmer Go-Lang terbiasa membuat table benchmark juga
- Ini digunakan agar kita bisa mudah melakukan performance test dengan kombinasi data berbeda-beda tanpa harus membuat banyak benchmark function

```go
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Burhanudin",
			request: "Burhanudin",
		},
		{
			name:    "Rabbani",
			request: "Rabbani",
		},
		{
			name:    "Burhanudin Rabbani",
			request: "Burhanudin Rabbani",
		},
		{
			name:    "Budi Nugraha",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}
```
