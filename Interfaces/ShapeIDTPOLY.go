package main

import (
	"fmt"
)

type circle struct {
	radius float64
}

type Rectangle struct {
	width, length float64
}

type shape interface {
	area() float64
	perimeter() float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.length
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.length)
}

func print(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter", s.perimeter())

}

func main() {
	// Interface Dynamic Type and Polymorphism
	var s shape
	fmt.Printf("Type: %T \n", s)
	ball := circle{radius: 2.5}
	s = ball
	print(s)

	//room := Rectangle{width: 45, length: 30}
	//s = room
	//print(s)

	//failed type assertion
	var sha shape = Rectangle{8, 10}
	sha.(Rectangle).perimeter()

	// type Assertion
	r, ok := sha.(Rectangle)
	if ok == true {
		fmt.Println("Rectangle perimeter \n", r.perimeter())
	}

	// Type Switch

	switch value := s.(type) {
	case circle:
		fmt.Println("%#v Circle type", value)
	case Rectangle:
		fmt.Println("%#V rectangle type", value)
	}

}
