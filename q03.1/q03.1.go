package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 引数の個数をチェック
	length := len(os.Args)
	if length != 3 {
		fmt.Print("Please set 2 arg\n")
		os.Exit(1)
	}

	baseFileName := os.Args[1]
	newFileName := os.Args[2]

	// コピー元ファイルを開く
	reader, err := os.Open(baseFileName)
	if err != nil {
		fmt.Printf("Failed to open file named: %v\n", baseFileName)
	}

	// 新規ファイルを作成
	file, err := os.Create(newFileName)
	if err != nil {
		fmt.Println("Failed to Create file")
		fmt.Println(err)
	}

	io.Copy(file, reader)

	fmt.Printf("コピー元: %v\nコピー先: %v\n", baseFileName, newFileName)
}
