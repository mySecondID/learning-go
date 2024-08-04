package main

import (
	custom1 "myPackage/custom"
)

func main() {
	custom1.Setval(12)
	custom1.Val = 122
	custom1.PrintVal()
}
