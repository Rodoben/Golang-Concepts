package main

import "fmt"

//  type emptyInterface interface {

//  }

type person struct {
	info interface{}
}

func main() {

	var empty interface{}
	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{8, 5, 6, 3, 5, 2}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "your name"
	fmt.Println(you.info)
	you.info = 40
	fmt.Println(you.info)
	you.info = []int{5, 6, 8, 7, 9, 3}
	fmt.Println(you.info)

}
