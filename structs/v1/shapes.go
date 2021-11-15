package shapes

import "math"

type Shape interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width * r.Height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
