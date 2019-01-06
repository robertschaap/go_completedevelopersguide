package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
