package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("This is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World\n")
	_, _ = writer.WriteString("Selamat belajar\n")
	_, _ = writer.WriteString("Semangaattt ❤️‍🔥❤️‍🔥❤️‍🔥\n")

	writer.Flush()
}
