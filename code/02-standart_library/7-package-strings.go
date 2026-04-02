package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Burhanudin Rabbani", "bani"))    // true
	fmt.Println(strings.Split("Burhanudin Rabbani Aurelius", " ")) // []string
	fmt.Println(strings.ToLower("Burhanudin Rabbani Aurelius"))
	fmt.Println(strings.ToUpper("Burhanudin Rabbani Aurelius"))
	fmt.Println(strings.Trim("            Burhanudin Rabbani Aurelius             ", " "))
	fmt.Println(strings.ReplaceAll("Burhanudin Rabbani Aurelius bani", "bani", "respect"))

}
