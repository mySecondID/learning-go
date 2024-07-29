package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("worker %d starts\n", i)

	// dummy operations are carried out

	time.Sleep(time.Second)

	fmt.Printf("worker %d ends\n", i)

}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerNo int) {
			defer wg.Done()
			worker(workerNo)
		}(i)
	}

	wg.Wait()

	fmt.Printf("%d\n", (time.Now().Nanosecond()-startTime.Nanosecond())/1000)
}
