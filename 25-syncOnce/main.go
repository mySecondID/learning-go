package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	// var mutexLock sync.Mutex

	var counter int32 = 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			// mutexLock.Lock()
			// counter++
			// mutexLock.Unlock()
			atomic.AddInt32(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("%d\n", counter)
}

// go run -race main.go
