package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Namaskar bhai"
	fmt.Println(welcome)

	read := bufio.NewReader(os.Stdin)

	// instead of try-catch, use comma-ok or err-ok
	input, _ := read.ReadString('\n')
	fmt.Println(input)

}
