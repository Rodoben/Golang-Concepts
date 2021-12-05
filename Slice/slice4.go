package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Please run the program with more than two nums")
	}

	args := os.Args[1:]
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Plese enter between 2 and 10 numbers")

	} else {
		for _, r := range args {
			num, err := strconv.ParseFloat(r, 64)
			if err != nil {
				continue

			}
			sum += num
			product *= num
		}
	}
	fmt.Printf("Sum: %v, Product: %v\n", sum, product)
}
