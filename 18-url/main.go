package main

import (
	"fmt"
	"net/url"
)

func main() {
	const URL string = "https://lco.dev:3000/learn?paymentid=123123&coursename=golang"
	parsedURL, _ := url.Parse(URL)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Port())
	fmt.Println(parsedURL.RawQuery)
	fmt.Println(parsedURL.Path)
	params := parsedURL.Query()
	fmt.Println(params)

	khudKaURL := &url.URL{
		Host:     "lco.dev",
		Scheme:   "https",
		Path:     "learn",
		RawQuery: "paymentid=123123&coursename=golang",
	}
	fmt.Println(khudKaURL)
}
