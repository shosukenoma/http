package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	// create a slice of type byte with 99999 empty elements.
	// this is because read() cannot resize []byte if it becomes full.
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}