package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.krebsonsecurity.com/feed/")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Pasrse body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
