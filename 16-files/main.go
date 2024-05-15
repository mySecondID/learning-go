package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readMyFile(filePath string) {
	input, err := os.ReadFile(filePath)
	// if err != nil {
	// 	panic(err)
	// }
	checkError(err)
	fmt.Println(string(input))
}

func main() {
	contentToBeShifted := "get lost"
	file, err := os.Create("./mygofile.txt")
	checkError(err)
	len, err := io.WriteString(file, contentToBeShifted)
	checkError(err)
	fmt.Println(len)
	readMyFile("./mygofile.txt")
}
