package main

import "time"

func Repeat(char string, count int) (result string) {
	for i := 0; i < count; i++ {
		result += char
	}
	time.Sleep(time.Millisecond * 100)
	return
}
