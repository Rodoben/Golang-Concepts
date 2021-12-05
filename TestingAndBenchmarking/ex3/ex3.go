package main

import "fmt"

func main() {
	fmt.Println("sum of 2+3:", Calculate(2, 3))
	fmt.Println("sum of 5+7+1+1:", Calculate(5, 7, 1, 1))
}

func Calculate(x ...int) int {
	sum := 0
	for _, r := range x {
		sum += r
	}
	return sum
}
