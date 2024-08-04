package main

import (
	"fmt"
	"sync"
	"time"
)

type SomeObj struct {
	data []byte
}

func makeSomeObj() *SomeObj {
	res := SomeObj{
		data: make([]byte, 1024*1024),
	}
	return &res
	// return &SomeObj{
	// 	data: make([]byte, 1024 * 1024),
	// }
}

func main() {
	// var objArr []*SomeObj
	var wg sync.WaitGroup
	var ctr int64 = 0

	objPool := sync.Pool{
		New: func() interface{} {
			ctr++
			return makeSomeObj()
		},
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			y := objPool.Get().(*SomeObj)
			// objArr = append(objArr, y)
			objPool.Put(y)
		}()
	}

	wg.Wait()
	time.Sleep(5 * time.Second)

	fmt.Println(ctr)
}
