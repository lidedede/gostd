package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) ShowMe(s string) string{
	return s
}

func main() {

	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	//var p People = s
	//if p == nil {
	//	fmt.Println("p is nil")
	//} else {
	//	fmt.Println("p is not nil")
	//}
	fmt.Println(s.ShowMe("test"))

}