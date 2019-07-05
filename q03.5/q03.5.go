package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world")
	copyN(os.Stdout, reader, 4)
}

func copyN(dest io.Writer, src io.Reader, length int64) {
	lReader := io.LimitReader(src, length)
	io.Copy(dest, lReader)
}
