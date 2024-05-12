package main

import (
	"fmt"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(input)
	// }

	// go env
	// GOOS="windows" go build
	currTime := time.Now()
	fmt.Println(currTime.Format("02-01-2006 Monday"))
}
