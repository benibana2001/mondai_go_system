package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("encodingCsv.csv")
	if err != nil {
		fmt.Println(err)
	}

	// CSVへのデコレータを介す
	writer := csv.NewWriter(file)
	record := []string{"encoding", "csv", "example"}
	writer.Write(record)
	writer.Flush()
}
