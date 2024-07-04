package structsmethodsandinterfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}

func Perimeter(width float64, height float64) float64 {
	return (width + height) * 2
}

func Perimeter2(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}

func Area(width float64, height float64) float64 {
	return width * height
}

func Area2(rect Rectangle) float64 {
	return rect.Width * rect.Height
}
