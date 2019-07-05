package main

import (
	"archive/zip"
	"io"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// io.Writer をメンバに持つ構造体 zipWriter を生成. zipの出力先はhttp.ResponseWriter
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	// zip化するファイルを作成し、その中身へのWrite() を持つ io.Writer を返す
	writer, _ := zipWriter.Create("new.txt")
	reader := strings.NewReader("Example of zipWriter\n")

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=new.zip")
	io.Copy(writer, reader)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8087", nil)
}

