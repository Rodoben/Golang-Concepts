package main

import "fmt"

func main() {

	// Anonymous function

	func(x ...int) {
		s := 0
		for _, c := range x {
			s = s + c
		}

		fmt.Println(s)
	}(10000, 2000)

	a := f1(10)
	for {

		if a() == 20 {
			break
		} else {
			fmt.Println(a())
		}

	}

}

func f1(x int) func() int {

	return func() int {
		x++
		return x
	}
}
