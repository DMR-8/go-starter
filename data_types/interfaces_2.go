package main

import (
	"fmt"
)

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(2.44)
	Anything("DMR-8")
	Anything(struct{}{})
	mymap := make(map[string]interface{})
	mymap["name"] = "DMR-8"
	mymap["age"] = 110
	fmt.Println(mymap)
}