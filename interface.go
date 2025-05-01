package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %v\n", s.Area())
}
