package main

import (
	"fmt"
)

func main() {
	initArr()
	appendArr()
}

//Init Arrays
func initArr() {
	arr := []int{1,2,3,4}
	arr2 := []string{"this", "is", "sparta"}
	fmt.Println(arr, arr2)
}

//Append Array
func appendArr(){
	arr2 := []string{"this", "is", "sparta"}
	arr2 = append(arr2, "!!!")
	fmt.Println(arr2)
}
