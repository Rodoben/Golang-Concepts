package main

import "fmt"

type product struct {
	price float64
	name  string
}

func ChangeProduct(p product) {
	p.price = 300
	p.name = "Bicycle"
}

func ChangeProductByPointer(p *product) {
	(*p).price = 300
	(*p).name = "Bicycle"
}

func main() {

	gift := product{price: 100, name: "car"}
	fmt.Println("Before Change Product:", gift)
	ChangeProduct(gift)
	fmt.Println("After Change Product:", gift)
	fmt.Println("******************************")
	fmt.Println("Before Change Product By Pointer:", gift)
	ChangeProductByPointer(&gift)
	fmt.Println("After Change Product By Pointer", gift)

}
