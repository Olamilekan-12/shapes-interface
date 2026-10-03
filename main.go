package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() float64 {
	return 0
}

func describe(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Println("circle with radius", v.Radius)
	case Rectangle:
		fmt.Println("rectangle", v.Width, "x", v.Height)
	default:
		fmt.Println("unknown shape")
	}
}

func main() {
	circle := Circle{
		Radius: 5,
	}

	rectangle := Rectangle{
		Width:  4,
		Height: 6,
	}

	triangle := Triangle{
		Base:   5,
		Height: 10,
	}

	shapes := []Shape{
		circle,
		rectangle,
		triangle,
	}

	for _, s := range shapes {
		describe(s)
	}

}
