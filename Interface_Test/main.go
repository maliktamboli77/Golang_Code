package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func calculateArea(s Shape) float64 {
	return s.Area()
}

func main() {
	rect := Rectangle{width: 7, height: 7}
	circle := Circle{radius: 4}

	fmt.Println("Rectangle Area: ", calculateArea(rect))
	fmt.Println("Circle Area: ", calculateArea(circle))
}
