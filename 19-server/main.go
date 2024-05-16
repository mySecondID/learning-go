package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var responseString strings.Builder
	cookie, _ := os.ReadFile("./.env")
	responseString.Write(cookie)
	fmt.Println(responseString.String())
	x := responseString.String()
	reqBody := strings.NewReader(fmt.Sprintf(`
		{
			"token": "%s" 
		}
	`, x))
	fmt.Println(reqBody)
}
