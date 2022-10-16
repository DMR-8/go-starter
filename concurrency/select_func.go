package main

import (
	"fmt"
	"os"
	"time"
)


func Select(c chan int, quits chan struct{}) {
	// switch case
	// select for channel async
	for {
		time.Sleep(time.Second)
		select { // switch case on channels
			case <-c:
				fmt.Println("received")
				quits <- struct{}{}
			case <-quits:
				fmt.Println("Quit")
				os.Exit(0)
		}
	}
}

func main() {
	c := make(chan int, 2)
	quits := make(chan struct{})

	go func() {
		c <- 1
	}()

	go Select(c, quits)
	select {}
}