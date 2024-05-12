package main

import "fmt"

func main() {
	var list [4]string
	list[0] = "apple"
	list[2] = "apple1"
	list[3] = "apple2"
	fmt.Println(list)
	fmt.Println(len(list))
	var _sec = [3]string{"yo", "yoyo", "yoyoyo"}
	fmt.Println(_sec)
}
