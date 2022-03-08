package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	str := `{"name":"lidedong","age":25}`
	//i := struct {
	//	make    string
	//	model   string
	//	mileage int
	//}{
	//	make:    "Ford",
	//	model:   "Taurus",
	//	mileage: 200000,
	//}
	t := reflect.TypeOf(str).String()
	fmt.Println(t)
	bytes, err := json.Marshal(str)
	if err != nil {
		fmt.Println("----1---", "item Marshal: %v", err)
	}
	fmt.Println("----2---", string(bytes))
}
