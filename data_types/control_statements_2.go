package main

import (
	"fmt"
)

func main() {

	flag := true
	// continue, break, switch case,
	for i := 1; i < 10; i++ {
		if i%4 == 0 {
			flag = false
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println(flag)

	day := "Fri"
	switch day {
		case "Fri":
			fmt.Println("TGIF")
		case "Mon", "Tue", "Wed":
			fmt.Println("Boring")
			break
		default:
			fmt.Println("default")
	}

	switch {
		case day == "Fri":
			fmt.Println("Second TGIF")
	}
}