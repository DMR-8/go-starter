package main

import "fmt"

func main() {
	declare()
	declareMultiple()
	typeCasting()
	initVariables()
}
// Declaring a variable 
func declare() {
	var m1 int
	m1=2
	fmt.Println(m1)
}
// Declaring multiple variable 
func declareMultiple() {
	var(
		m1 = 2
		m2 = 3
	)
	fmt.Println(m1+m2)
}
// Declaring multiple variable 
func typeCasting() {
	var m1 int32
	var m2 int64
	fmt.Println(int64(m1)+m2)
}
// Initiali variables
func initVariables() {
	m1 := 2
	m2 := 3
	fmt.Println(m1+m2)
}