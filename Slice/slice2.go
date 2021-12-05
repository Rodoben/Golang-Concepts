package main

import "fmt"

func main() {
	a1 := []string{"ronald", "foad", "linc"}

	for i, r := range a1 {
		fmt.Printf("Index %d  Value %s \n", i, r)
	}

	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6

	a := 10.5
	mySlice[0] = a

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

}
