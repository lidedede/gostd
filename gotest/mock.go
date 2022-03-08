package main

import "fmt"

func Add(a,b int) int {
	fmt.Println("add")
	return a+b
}

func NeedMock(a int) int {
	sum := Add(3,2)
	return sum + a
}

