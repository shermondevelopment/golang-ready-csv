package main

import (
	"encoding/csv"
	"fmt"
	"os"
)


func main() {
	println("hi, i'm go")

	/* open file */
	file, error := os.Open("clients.csv")

	if error != nil {
		panic(error)
	}

	readyCsv := csv.NewReader(file)

	for {
		record, error := readyCsv.Read()

		if error != nil {
			break
		}
		fmt.Println(record)
	}
}