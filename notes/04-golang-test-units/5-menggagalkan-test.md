# Menggagalkan test

- Menggagalkan unit test menggunakan panic bukanlah hal yang bagus
- Go-Lang sendiri sudah menyediakan cara untuk menggagalkan unit test menggunakan testing.T
- Terdapat function Fail(), FailNow(), Error() dan Fatal() jika kita ingin menggagalkan unit test

## t.Fail() dan t.FailNow()

- Terdapat dua function untuk menggagalkan unit test, yaitu Fail() dan FailNow(). Lantas apa bedanya?
- Fail() akan menggagalkan unit test, namun tetap melanjutkan eksekusi unit test. Namun diakhir ketika selesai, maka unit test tersebut dianggap gagal
- FailNow() akan menggagalkan unit test saat ini juga, tanpa melanjutkan eksekusi unit test

```go
func TestHelloWorldBani(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Fail() // Ini akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Ini test Hello World Bani")

}


func TestHelloWorldUdin(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.FailNow() // Ini tidak akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Ini test Hello World Udin")

}
```

## t.Error(args...) dan t.Fatal(args...)

- Selain Fail() dan FailNow(), ada juga Error() dan Fatal()
- Error() function lebih seperti melakukan log (print) error, namun setelah melakukan log error, dia akan secara otomatis memanggil function Fail(), sehingga mengakibatkan unit test dianggap gagal
- Namun karena hanya memanggil Fail(), artinya eksekusi unit test akan tetap berjalan sampai selesai
- Fatal() mirip dengan Error(), hanya saja, setelah melakukan log error, dia akan memanggil FailNow(), sehingga mengakibatkan eksekusi unit test berhenti

```go
func TestHelloWorldBani(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Error("Result must be Hello Bani") // Ini akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Dieksekusi walaupun Error")

}


func TestHelloWorldUdin(t *testing.T){
	result := HelloWorld("Bani")

	if result != "Hello Bani"{
		// unit test failed
		t.Fatal("Result must be Hello Bani") // Ini tidak akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Tidak dieksekusi ketika Error")

}
```

Next: [Assertions](./6-assertions.md)
