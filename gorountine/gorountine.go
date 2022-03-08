package main

import (
	"fmt"
	"time"
	//"runtime"
)
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1000 * time.Millisecond)

		fmt.Println(s)
	}
	//fmt.Println("-----------",s)
}
func main() {
	go say("world")
	time.Sleep(1000 * time.Millisecond)
	go say("world111")
	//runtime.GOMAXPROCS(2)
	say("hello")
}
