package main

import (
	"fmt"
)

func main() {
	initPointer()

	m1,m2 := 2,3 
	swap(&m1, &m2)// Pass by reference
	fmt.Println(m1, m2)

}

//Init Pointers
func initPointer() {
	m1 := 2 
	ptr := &m1
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

//Append Array
func swap( m1, m2 * int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}