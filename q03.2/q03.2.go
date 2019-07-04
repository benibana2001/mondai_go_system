package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func main() {
	// ランダムな文字で埋められたファイルを作成
	// ファイルサイズは 1024 byte とする

	writer, err := os.Create("new.txt")
	if err != nil {
		fmt.Println("Failed to create new file")
		os.Exit(1)
	}

	reader := rand.Reader
	lReader := io.LimitReader(reader, 1024)

	io.Copy(writer, lReader)
}
