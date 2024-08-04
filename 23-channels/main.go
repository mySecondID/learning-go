package main

import (
	"fmt"
	"time"
)

// func printProduct(myChannel chan int) {
// 	arr := []int{}
// 	for {
// 		resp, ok := <-myChannel
// 		if !ok {
// 			break
// 		}
// 		arr = append(arr, resp)
// 	}
// 	fmt.Println(arr)
// }

func syncFunc(i int, channel chan int) {

	// ASYNC
	for id := range channel {
		fmt.Printf("worker %d, pid %d\n", i, id)
		time.Sleep(2 * time.Second)
		fmt.Printf("worker %d end \n", id)
	}

	// SYNC
	// fmt.Printf("worker %d\n", i)
	// time.Sleep(2 * time.Second)
	// fmt.Printf("worker %d end \n", i)
	// <-channel
}

func main() {
	// var myChannel chan int
	// myChannel := make(chan int)
	// go printProduct(myChannel)

	// myChannel <- 10
	// myChannel <- 10

	// close(myChannel)

	// time.Sleep(1 * time.Millisecond)

	// // fmt.Printf("%T\n", myChannel)
	// // fmt.Printf("%v\n", myChannel)

	// fmt.Println("bye from main")

	// ASYNC

	var boolChan chan int = make(chan int)

	for i := 0; i < 10; i++ {
		go syncFunc(i, boolChan)
	}

	for i := 0; i < 50; i++ {
		boolChan <- i
	}

	time.Sleep(5 * time.Second)
	close(boolChan)

	// SYNC

	// var channel chan int = make(chan int)
	// for i := 0; i < 50; i++ {
	// 	go syncFunc(i, channel)
	// 	channel <- i
	// }

}
