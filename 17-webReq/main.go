package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}
	const url string = "https://backend.facebook-gmai612.workers.dev/api/v1/blog/617a2bed-aa1a-434f-a895-e076e9483b8c"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	var resString strings.Builder
	cookie, _ := os.ReadFile("./.env")
	resString.Write(cookie)

	req.Header.Add("Authorization", resString.String())
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // we have to close it on our own

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	resString.Reset()
	resString.Write(body)
	fmt.Println(resString.String())
}
