package main

import (
	"fmt"
)

type GreetingMessage struct {
	message string
}
type Area struct {
	radius int
	length int
	width  int
}

func (a Area) AreaCircle() float64 {
	area := (float64(a.radius * a.radius)) * 3.14
	return area
}

func (a Area) AreaRectangular() int {
	area := a.length * a.width
	return area
}

func (s GreetingMessage) PrintMessage() string {
	message := s.message
	return message
}

type Representer interface {
	AreaCircle() float64
	AreaRectangular() int
}

func main() {
	message := GreetingMessage{message: "Hello from other worlds"}
	fmt.Println(message.PrintMessage())

	circl := Area{radius: 10, length: 10, width: 4}
	areaCalculator(circl)
}

func areaCalculator(r Representer) {
	fmt.Println(r.AreaRectangular())
	fmt.Println(r.AreaCircle())
}
