package main

import (
	"encoding/csv"
	"os"
)

func main() {
	/*
		var value string = "Burhanudin Rabbani"
		var encoded = base64.StdEncoding.EncodeToString([]byte(value))
		fmt.Println(encoded)

		decoded, err := base64.StdEncoding.DecodeString(encoded)

		if err != nil {
			fmt.Println("Error:", err.Error())
			} else {
				fmt.Println(string(decoded))
			}
			}

			csvString := "Burhanduin,Rabbani,Aurelius\n" +
			"Ryan,Hidayat,Ganja\n" +
			"Heri,Pratama,Protein"

			reader := csv.NewReader(strings.NewReader(csvString))

			for {
				record, err := reader.Read()

				if err == io.EOF {
					break
				}

				fmt.Println(record)
	*/

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Burhanudin", "Rabbani", "Aurelius"})
	_ = writer.Write([]string{"Ryan", "Hidayat", "AI"})
	_ = writer.Write([]string{"Heri", "Pratama", "Protein"})

	writer.Flush()
}
