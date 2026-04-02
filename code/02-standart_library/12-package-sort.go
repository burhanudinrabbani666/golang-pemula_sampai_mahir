package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// Standard example Golang https://pkg.go.dev/sort@go1.26.1
func (s UserSlice) Len() int           { return len(s) }
func (s UserSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s UserSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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
