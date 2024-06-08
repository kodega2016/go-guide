package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	lw := logWriter{}
	// defer res.Body.Close()
	// bs := make([]byte, 99999)
	// res.Body.Read(bs)
	// fmt.Println("response:", string(bs))
	defer res.Body.Close()
	io.Copy(&lw, res.Body)
}

func (lw logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("just wrote %d bytes in the console\n", len(bs))
	return len(bs), nil
}
