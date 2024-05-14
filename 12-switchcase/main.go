package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	diceNo := rand.Intn(6) + 1
	switch diceNo {
	case 1:
		fmt.Println("1;;")
	case 2:
		fmt.Println("2;;")
	case 4:
		fmt.Println("4;;")
		fallthrough
	case 3:
		fmt.Println("3;;")
	default:
		fmt.Println("d;;")
	}
}
