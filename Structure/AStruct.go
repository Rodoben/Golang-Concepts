package main

import "fmt"

// anaonymous filed
type Book struct {
	string
	float64
	int
}

type Employee struct {
	name   string
	salary int
	bool
}

func main() {
	// anonymous struct
	ronald := struct {
		firstname, lastname string
		age                 int
	}{
		firstname: "Ronald",
		lastname:  "Benjamin",
		age:       30,
	}
	fmt.Println("Anonymous Struct", ronald)
	ronald.lastname = "George"
	ronald.firstname = ""
	fmt.Println("Anonymous Struct", ronald)

	b1 := Book{"Animal Farm", 444.44, 4}
	fmt.Println("Book:", b1)

	e2 := Employee{"Ronald", 33333, true}
	fmt.Println("Employee:", e2)

	e2.bool = false

	fmt.Println("Employee:", e2)

}
