package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

type geometry interface {

	// Embedding of two interface shape and object
	shape
	object
	getColor() string
}

type cube struct {
	edge float64
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volume() float64 {

	return math.Pow(c.edge, 3)

}

func (c cube) getColor() string {
	return "yellow"
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	b := g.volume()
	return a, b
}

func main() {
	c := cube{5.3}
	a, b := measure(c)
	fmt.Println(a, b)
}
