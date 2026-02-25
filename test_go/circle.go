package test_go

import "math"

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
