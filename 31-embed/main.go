package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed test.json
var data []byte

type Data struct {
	Phone string
	Name  string
}

func main() {
	fmt.Println(string(data))
	var res Data
	json.Unmarshal(data, &res)
	fmt.Println(res.Name)
	fmt.Println(res.Phone)
}
