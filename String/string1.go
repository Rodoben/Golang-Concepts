package main

import "fmt"

func main() {
	p := fmt.Println
	s1 := "Hi there Gopher"
	p(s1)

	fmt.Println("He says :\"Hello!\"")
	//fmt.Println( 'He says: "Hello!"')

	s2:= 'I like to code'
	fmt.Println(s2)
}
