package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {

	file, _ := os.Create("new.zip")
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, _ := zipWriter.Create("new.txt")
	reader := strings.NewReader("Example of zipWriter\n")

	io.Copy(writer, reader)
}
