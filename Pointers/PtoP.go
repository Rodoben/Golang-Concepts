package main

import "fmt"

func main() {
	a := 5
	p1 := &a
	pp1 := &p1

	fmt.Printf("Value of P1: %V, Address of P1:%V", p1, &p1)
	fmt.Println()
	fmt.Printf("Value of PP1: %V, Address of PP1:%V", pp1, &pp1)

	// dereference  a pointer to pointer
	fmt.Println()
	fmt.Printf("*p1 is %V \n", *p1)
	fmt.Printf("*pp1 is %V \n", *pp1)
	fmt.Println()
	fmt.Printf("**pp1 is %V \n", **pp1)
	**pp1++
	fmt.Printf("a is %v \n", a)

	var p2 *int
	fmt.Printf("%#V \n", p2)
	var p3 *int

	y, z := 5, 5
	p2 = &y
	p3 = &z
	fmt.Println(y == z)
	fmt.Println(p2)
	fmt.Println(p3)
	fmt.Println(*p2 == *p3)

}
