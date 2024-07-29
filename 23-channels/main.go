package main

import (
	"fmt"
	"time"
)

func printProduct(myChannel chan int) {
	arr := []int{}
	for {
		resp, ok := <-myChannel
		if !ok {
			break
		}
		arr = append(arr, resp)
	}
	fmt.Println(arr)
}

func main() {
	// var myChannel chan int
	myChannel := make(chan int)
	go printProduct(myChannel)

	myChannel <- 10
	myChannel <- 10

	close(myChannel)

	time.Sleep(1 * time.Millisecond)

	// fmt.Printf("%T\n", myChannel)
	// fmt.Printf("%v\n", myChannel)

	fmt.Println("bye from main")
}
