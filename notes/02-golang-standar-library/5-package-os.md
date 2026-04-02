# Package os

- Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
- Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan disemua sistem operasi)
- https://golang.org/pkg/os/

```go
func main() {
	// For get argument in command line
	args := os.Args

	for _, arg := range args {
		fmt.Println(arg)
	}

}
```

```go
func main() {
	hostname, err := os.Hostname()

	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Hostname:", hostname)
	}
}
```

Next: [Package flag](./6-package-flag.md)
