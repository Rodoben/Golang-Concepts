package main

import "fmt"

// Embedded Struct
type contact struct {
	email, address string
	phno           int
}

type Employee struct {
	name  string
	age   int
	cInfo contact // Embedded.

}

func main() {
	john := Employee{
		name: "Ronald Benjamin",
		age:  24,
		cInfo: contact{
			email:   "dodo.ben.rb@gmail.com",
			address: "belabagan",
			phno:    998639,
		},
	}

	fmt.Println(john)
	john.age = 56
	john.cInfo.phno = 333
	fmt.Println(john)

}
