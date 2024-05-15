package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	cookie, _ := os.ReadFile("./.env")
	client := &http.Client{}
	const url string = "https://backend.facebook-gmai612.workers.dev/api/v1/blog/617a2bed-aa1a-434f-a895-e076e9483b8c"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", string(cookie))
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // we have to close it on our own

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
