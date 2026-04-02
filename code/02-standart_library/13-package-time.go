package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		// Set time now
		var timeNow time.Time = time.Now()
		fmt.Println(timeNow.Local())

		// Set manually
		utc := time.Date(2002, time.November, 14, 23, 0, 0, 0, time.UTC)
		fmt.Println(utc.Local())

		// Using Format
		formater := "2006-01-02 15:04:05"
		value := "2020-10-10 10:10:10"
		valueTime, err := time.Parse(formater, value)

		if err != nil {
			fmt.Println("Error", err.Error())
			} else {
				fmt.Println(valueTime)
			}

			fmt.Println(timeNow.Year())
			fmt.Println(timeNow.Month())
			fmt.Println(timeNow.Date())
			fmt.Println(timeNow.Day())
			fmt.Println(timeNow.Hour())
	*/

	var duration1 time.Duration = 100 * time.Second
	var duration2 time.Duration = 10 * time.Millisecond
	var duration3 time.Duration = duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("duration: %d\n", duration3)

}
