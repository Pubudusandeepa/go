package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

// Circle define circle
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return circle.x * circle.y * math.Pi
}

func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
