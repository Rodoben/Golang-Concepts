package main

import "fmt"

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

// func (r Rectangle) perimeter() float64 {
// 	return 2 * (r.width + r.length)
// }

func print(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())

}
func main() {

	c1 := circle{2.3}
	//fmt.Println(c1.area())
	print(c1)
	r1 := Rectangle{25, 20}
	print(r1)

}
