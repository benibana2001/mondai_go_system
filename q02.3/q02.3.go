package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Set Header
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content_Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	file, err := os.Create("json.txt.gzip")
	if err != nil {
		fmt.Println(err)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")


	writerGzip := gzip.NewWriter(file)
	writerGzip.Header.Name = "json.txt"
	writerGzip.Close()
	//result, _ := json.Marshal(&source)
	//io.WriteString(writerGzip, string(result))

	// セットしたio.Writer にエンコードと書き出しを行う
	//encoder.Encode(source)
	//io.WriteString(w, "http.Response Writer sample")
}
