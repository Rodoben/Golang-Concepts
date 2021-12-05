package main

import "fmt"

func main() {

	x := 8
	p := &x
	fmt.Println("Change Func before passing value", x, p)
	//change(p)
	fmt.Println("Change Func After passing value", x)

	fmt.Println("************************************")
	fmt.Println("Change Func before passing value", x)
	changeVar(x)
	fmt.Println("Change Func After passing value", x)
}

func change(a *int) *float64 {
	*a = 100
	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 88
}
