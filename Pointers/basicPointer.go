package main

import "fmt"

func main() {

	// To print the address of variable name
	name := "Ronald"
	fmt.Println("Address of name", &name)
	//*********************************************
	// Memory Address of x
	var x int = 2
	ptr := &x
	fmt.Printf("%T %P", ptr, ptr)
	fmt.Println("\n")
	fmt.Printf("%P", &x)

	//*********************************************
	var ptr1 *float64
	_ = ptr1
	fmt.Println("\n")
	fmt.Printf("ptr1 %P", ptr1)

	p := new(int)
	fmt.Println("\n")
	fmt.Println("new p", p)
	fmt.Println("\n")
	x = 100
	p = &x
	fmt.Printf("P is of type %T with a value of %V \n", p, p)

	//*p = 90
	fmt.Println(*p)

	fmt.Println("*p==x:", *p == x)

	*p = *p / 2
	fmt.Println("X value", x)
}
