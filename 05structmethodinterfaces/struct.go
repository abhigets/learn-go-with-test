package _5structmethodinterfaces

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}

type Shape interface {
	Area() float64
}

type Triangle struct {
	Length float64
	Height float64
}

func (triangle Triangle) Area() float64 {
	return .5 * (triangle.Height * triangle.Length)
}
