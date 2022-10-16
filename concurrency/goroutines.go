package main

import (
	"fmt"
	"time"
)
func heavy() {
	for { //infinite for loop
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}
func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super heavy")
	}
}
func main() {
	go heavy() // goroutines
	go superHeavy()
	fmt.Println("Fin")
	select {} // waits indefinitely
}