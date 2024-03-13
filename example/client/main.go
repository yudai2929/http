package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body := make([]byte, 1024)
	resp.Body.Read(body)

	fmt.Println(string(body))
}
