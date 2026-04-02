# Package flag

- Package flag berisikan fungsionalitas untuk memparsing command line argument
- https://golang.org/pkg/flag/

```go
import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database username")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Printf("Username: %s \n", *username)
	fmt.Printf("Password: %s \n", *password)
	fmt.Printf("host: %s \n", *host)
	fmt.Printf("Port: %d \n", *port)
}
```

Contoh dibash

```bash
go run 6-package-flag.go -username=bani -password=secret -host=123.45.678 -port=8080
```

Next: [Package strings](./7-package-strings.md)
