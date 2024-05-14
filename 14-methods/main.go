package main

import "fmt"

type User struct {
	Name     string
	Email    string
	password string
	cntr     int64
}

func (u *User) IncAndPrintbyRef() {
	fmt.Println(u.cntr)
	u.cntr++
}

func (u User) IncAndPrint() {
	fmt.Println(u.cntr)
	u.cntr++
}

func main() {
	a := new(User)
	var ptr *User = a
	ptr.IncAndPrint()
	ptr.IncAndPrint()
	ptr.IncAndPrint()
	ptr.IncAndPrint()
	ptr.IncAndPrint()

	ptr.IncAndPrintbyRef()
	ptr.IncAndPrintbyRef()
	ptr.IncAndPrintbyRef()
	ptr.IncAndPrintbyRef()
	ptr.IncAndPrintbyRef()
}
