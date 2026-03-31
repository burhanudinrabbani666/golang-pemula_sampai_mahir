package main

import "fmt"

func main() {

	data := newMap("Bani")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}

}

func newMap(name string) map[string]string {

	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
