package main

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// Rectangle model
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculation
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculation for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle model
type Circle struct {
	Radius float64
}

// Area calculation for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle model
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculation for Triangle
func (t Triangle) Area() float64 {
	// return Rectangle{t.Base, t.Height}.Area() * 0.5

	return (t.Base * t.Height) * 0.5
}
