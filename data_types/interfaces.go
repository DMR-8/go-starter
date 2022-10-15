package main

import (
	"fmt"
)
type Car interface {
	Drive()
	Stop()
}
type Lambo struct {
	model string
}
func NewModel(arg string) Car {
	return &Lambo{arg} //enforces the interface on Lambo struct
}
type Chevy struct {
	model string
}
func (l *Lambo) Stop() {
	fmt.Println("Stopping lambo")
}
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.model)
}
func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.model)
}
func main() {
	l := NewModel("Gallardo") // Lambo needs to find the interface
	c := Chevy{"C369"} // Chevy does not need to follow the interface
	l.Drive()
	c.Drive()
}