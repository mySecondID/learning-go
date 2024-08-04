package main

import "fmt"

type calcCGPA interface {
	update(semester int, marks int)
	display()
}

type Student struct {
	name  string
	email string
	marks int
}

func (s *Student) update(semester int, marks int) {
	(*s).marks += marks
}

func (s *Student) display() {
	fmt.Println((*s).marks)
}

func main() {
	var Pratham calcCGPA = &Student{"pratham", "abc@gmail.com", 0}
	var ptr *calcCGPA = &Pratham
	(*ptr).update(2, 10)
	(*ptr).update(2, 10)
	(*ptr).update(2, 10)
	(*ptr).update(2, 10)
	(*ptr).display()
	Pratham.display()
}
