package main

import "fmt"

func MyFunc() string {
	return "hi from myfunc"
}

func myMathFunc(a int, b int) int {
	if a < 0 {
		a = -1 * a
	}
	return a + b
}

func main() {
	fmt.Println(MyFunc())
	fmt.Println(myMathFunc(2, 3))
}
