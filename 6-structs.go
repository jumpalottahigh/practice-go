package main

import "fmt"

func main() {
	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width: 10}
	fmt.Println("Area of rectangle =", rect1.area())
}

type Rectangle struct {
	leftX float64
	topY float64
	height float64
	width float64
}

// Assign method to struct
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
