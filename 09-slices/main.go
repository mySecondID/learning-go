package main

import "fmt"

func main() {
	var list = []string{"a", "b", "c"}
	fmt.Printf("%T\n", list)
	list = append(list[1:3], "d", "e")
	fmt.Println(list)

	l2 := make([]int, 4)
	l2[0] = 1
	l2[1] = 2
	l2[2] = 3
	l2[3] = 4
	fmt.Println(&l2[0])
	l2 = append(l2, 5, 6, 7)
	fmt.Println(l2)
	fmt.Println(&l2[0])

	l2 = append(l2[:3], l2[4:]...)
	fmt.Println(l2)
}
