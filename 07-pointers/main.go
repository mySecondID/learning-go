package main

import "fmt"

// pointers are being used to pass values of arrays like DS by reference

func main() {
	var a int = 2
	var ptr *int = &a
	fmt.Println(ptr)
}
