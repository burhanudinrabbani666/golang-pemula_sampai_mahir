# Skip test

- Kadang dalam keadaan tertentu, kita ingin membatalkan eksekusi unit test
- Di Go-Lang juga kita bisa membatalkan eksekusi unit test jika kita mau
- Untuk membatalkan unit test kita bisa menggunakan function Skip()

```go
func TestSkio(t *testing.T){


	if runtime.GOOS == "linux"{
		t.Skip("Cant run on mac OS")
	}


	result := HelloWorld("Bani")
	require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

}
```

Next: [Before after test](./8-before-after-test.md)
