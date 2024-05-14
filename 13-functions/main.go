package main

import (
	"fmt"
)

func f1(s1 string, s2 string) int {
	s3 := s1[0] + s2[0]
	return int(s3)
}

func f2(b int, values ...int) int {
	s := 0
	for _, val := range values {
		s += val
	}
	return s - b
}

func main() {
	fmt.Println(f1("abc", "def"))
	fmt.Println(f2(1, 2, 3, 4, 5))
}
