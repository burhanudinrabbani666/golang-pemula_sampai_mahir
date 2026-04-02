# Package sort

- Package sort adalah package yang berisikan utilitas untuk proses pengurutan
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
- https://golang.org/pkg/sort/

```go
type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int { return len(s) }
func (s UserSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s UserSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func main() {
	users := []User{
		{"Bani", 23},
		{"Nico", 25},
		{"Vito", 15},
		{"Anang", 26},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
```

Next: [Package time](./13-package-time.md)
