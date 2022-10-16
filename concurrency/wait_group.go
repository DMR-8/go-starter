package main

import (
	"fmt"
	"sync"
)

//we can also use Mutex Locks instead of waitgroups

func main() {
	// wait groups
	wg := &sync.WaitGroup{}
	// add, done and wait
	wg.Add(2)
	go func() { // lambda/anonymous funciton
		fmt.Println("Hello")
		wg.Done()
	}() // How to call a lambda functions
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("Fin")
	wg.Wait() // waits ffor two signals to waitGroup
	fmt.Println("Exit")
}