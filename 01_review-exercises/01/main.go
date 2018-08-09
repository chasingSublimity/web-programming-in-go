package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

type shape interface {
	// method name, return type
	area() float64
}

func info(s shape) {
	fmt.Println("area: ", s.area())
}

func main() {
	c := circle{50}
	s := square{50}
	info(c)
	info(s)
}
