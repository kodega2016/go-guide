package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println("reading file:", filename)
	bs, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, bs)
}
