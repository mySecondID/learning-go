package main

import "fmt"

type User struct {
	Name     string
	email    string
	phone_no int
}

func main() {
	shashwat := User{"shashu", "abc", 123}
	fmt.Println(shashwat.Name)
}
