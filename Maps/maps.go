package main

import "fmt"

func main() {

	names := map[int]string{
		1: "ronald",
		2: "benjamin",
		3: "tanvi",
		4: "lucy",
	}

	city := map[string]string{
		"bangalore": "margondonahalli",
		"mumbai":    "thane",
		"kolkata":   "park street",
	}

	fmt.Println(names)
	delete(names, 1)
	fmt.Println(names)

	for key, value := range names {
		fmt.Println("The name, ", value, "and its key is", key)
	}

	PrintMap(names)
	PrintMapCity(city)

}

//iteration over the map using functions

func PrintMap(c map[int]string) {
	for key, value := range c {
		fmt.Println("The key is", key, "and its value is", value)
	}

}

func PrintMapCity(c map[string]string) {
	for key, value := range c {
		fmt.Println("The key is", key, "and its value is", value)
	}

}
