package main

import "fmt"

func main() {
	var s1 string
	s1 = "Learning go!"
	fmt.Println(s1)

	var k int = 6
	var i = 5
	var j = 5.6

	fmt.Println("i:", i, "j:", j, "K:", k)

	var s2 = "Go!"
	_ = s2

	var ii, jj int
	ii, jj = 5, 8
	ii, jj = jj, ii
	fmt.Println(ii, jj)

	s3 := "Learning golang!"
	_ = s3

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	var deleted = false
	deleted, file := true, "a.txt"
	_, _ = deleted, file

	sum := 5 + 2.5
	fmt.Println(sum)

	var (
		age       float64
		firstName string
		gender    bool
	)
	_, _, _ = age, firstName, gender

	var a, b int
	_, _ = a, b
}
