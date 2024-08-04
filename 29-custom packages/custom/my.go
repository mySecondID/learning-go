package custom1

import "fmt"

var (
	Val int32
	val int32
)

func Setval(i int32) {
	val = i
}

func PrintVal() {
	fmt.Printf("Val:%d val:%d\n", Val, val)
}
