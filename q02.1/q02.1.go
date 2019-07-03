package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	d := 123456
	f := 1.2345
	s := "12345"

	d2s := strconv.Itoa(d)
	f2s := fmt.Sprint(f)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}

	str := d2s + "\n" + f2s + "\n" + s + "\n"

	file.Write([]byte(str))
	file.Close()

	fmt.Fprintf(os.Stdout, str)
}
