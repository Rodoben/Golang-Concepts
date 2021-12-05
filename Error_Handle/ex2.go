package main

import (
	"errors"
	"fmt"
)

// We can create a simple custom error in Golang with the following syntax:

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

func main() {
	r, err := calculateArea(11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Area:", r)
}
