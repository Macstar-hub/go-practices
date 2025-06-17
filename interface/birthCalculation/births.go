package main

import (
	"fmt"
)

type Date struct {
	christian int
	solar     int
}

func (c Date) christianCalc() int {
	age := c.christian - 1996
	return age
}

func (s Date) solarCalc() int {
	age := s.solar - 1374
	return age
}

type calcAge interface {
	solarCalc() int
	christianCalc() int
}

func main() {
	age := Date{solar: 1404, christian: 2025}
	calculatorAge(age)

}

func calculatorAge(c calcAge) {
	fmt.Println("Christian age:", c.christianCalc())
	fmt.Println("Solar age:", c.solarCalc())
}
