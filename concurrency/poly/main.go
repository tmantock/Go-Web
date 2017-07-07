package main

import (
	"fmt"
	"math"
)

type Square struct {
	side int
}

type Circle struct {
	radius int
}

func (s Square) Area() int {
	return s.side * s.side
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(2, float64(c.radius))
}

func main() {
	c := Circle{4}
	s := Square{4}

	x := c.Area()
	y := s.Area()

	fmt.Println(x)
	fmt.Println(y)
}
