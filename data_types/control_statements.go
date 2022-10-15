package main

import (
	"fmt"
)

func main() {
	// if else, for, switch case, break continue
	f := true
	flag := &f
	if flag == nil { // Use && or || for chaining
		fmt.Println("Value is nil") // checkc if value is being initialized or not
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	arr := []string{"DMR-8", "Daio", "middlename"}
	for i, value := range arr {
		fmt.Println(i, value)
	}
	
	mymap := make(map[string]interface{})
	mymap["name"] = "DMR-8"
	mymap["age"] = 1100

	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v\n", k, v)
	}
}