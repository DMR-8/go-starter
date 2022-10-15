package main

import (
	"fmt"
	"strings"
)

func main() {
	mutableProp()
	checkContains()
	checkReplace()
	delimSplit()
	concat()
}

//Mutable Strings
func mutableProp() {
	m1 := "My Age"
	m1 = " Is 100" // We can only change the value after assigning once
	fmt.Println(m1)
}

//Contains Fucntion
func checkContains(){
	m1 := "My Age"
	m2 := "Age"
	fmt.Println(strings.Contains(m1,m2))
}

//Replace in String 
func checkReplace(){
	m1 := "my Age Name"
	fmt.Println(strings.ReplaceAll(m1,"m","NULL"))
}

//Split Strings 
func delimSplit(){
	m1 := "my Age Name"
	fmt.Println(strings.Split(m1," "))
}

//Concatenate
func concat(){
	m1 := "My Age"
	m2 := " Is 110"
	fmt.Println(m1+m2)
}