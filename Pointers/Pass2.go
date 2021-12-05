package main

import "fmt"

func main() {

	quantity, price, name, sold := 5, 300.4, "Ronald", true
	fmt.Println("Before calling ChangeValues()", quantity, price, name, sold)
	ChangeValues(quantity, price, name, sold)
	fmt.Println("After calling ChangeValues()", quantity, price, name, sold)
	fmt.Println("*********************************************************")
	fmt.Println("Before calling ChangeValuesByPointer()", quantity, price, name, sold)
	ChangeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After calling ChangeValuesByPointer()", quantity, price, name, sold)
}

func ChangeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func ChangeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}
