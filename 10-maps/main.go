package main

import "fmt"

func main() {
	langs := make(map[string]string)

	langs["JS"] = "javascript"
	langs["PY"] = "Python"
	langs["CPP"] = "C++"

	fmt.Println(langs)

	delete(langs, "JS")
	fmt.Println(langs)

	for i, val := range langs {
		fmt.Println(i, val)
	}
}
