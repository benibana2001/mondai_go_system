package main

import (
	"io"
	"os"
	"strings"
)

var (
	computer = strings.NewReader("COMPUTER")
	system = strings.NewReader("SYSTEM")
	programming = strings.NewReader("PROGRAMMING")
)

func main() {
	var stream io.Reader
	////
	A := io.NewSectionReader(programming, 5,1)
	S := io.NewSectionReader(system, 0, 1)
	C := io.NewSectionReader(computer, 0, 1)
	I := io.NewSectionReader(programming, 8, 1)
	I2 := io.NewSectionReader(programming, 8, 1)
	stream = io.MultiReader(A, S, C, I, I2)
	////
	io.Copy(os.Stdout, stream)
}
