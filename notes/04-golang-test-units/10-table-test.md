# Table test

- Sebelumnya kita sudah belajar tentang sub test
- Jika diperhatikan, sebenarnya dengan sub test, kita bisa membuat test secara dinamis
- Dan fitur sub test ini, biasa digunaka oleh programmer Go-Lang untuk membuat test dengan konsep table test
- Table test yaitu dimana kita menyediakan data beruba slice yang berisi parameter dan ekspektasi hasil dari unit test
- Lalu slice tersebut kita iterasi menggunakan sub test

```go
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld Bani",
			request:  "Bani",
			expected: "Hello Bani",
		},
		{
			name:     "HelloWorld Rian",
			request:  "Rian",
			expected: "Hello Rian",
		},
		{
			name:     "HelloWorld Heri",
			request:  "Heri",
			expected: "Hello Heri",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}

}
```

Next: [Mock](./11-mock.md)
