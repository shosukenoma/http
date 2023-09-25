package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// // create a slice of type byte with 99999 empty elements.
	// // this is because read() cannot resize []byte if it becomes full.
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	// // os.Stdout implements the Writer interface
	// // resp.Body implements the Reader interface

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Finished writing", len(bs), "bytes")

	return len(bs), nil
}