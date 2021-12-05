package main

import "fmt"

func main() {
	s1 := []float64{}
	n := 0.
	c, ch := 1, 1
	sum, product := 0.0, 1.0

	for c == 1 {
		fmt.Println("Enter a number")
		fmt.Scanln(&n)
		if n > 2 && n < 10 {
			s1 = append(s1, n)
		} else {
			fmt.Println("Plz enter no b/w 2 and 10")
		}
		fmt.Println("Do you want to continue enter 1: yes 2: No")
		fmt.Scanln(&ch)
		c = ch
	}
	fmt.Println(s1)
	for _, r := range s1 {
		sum += r
		product *= r
	}
	fmt.Println("Sum:", sum, "Product:", product)
}
