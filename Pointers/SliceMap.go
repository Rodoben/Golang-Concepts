package main

import "fmt"

func ChangeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func ChangeMap(m map[string]int) {
	m["a"] = 5
	m["b"] = 5
	m["c"] = 5
}

func ChangeMapByPointer(m *map[string]int) {
	(*m)["a"] = 10
	(*m)["b"] = 10
	(*m)["c"] = 10
}

func main() {
	prices := []int{1, 2, 3}
	ChangeSlice(prices)
	fmt.Println("prices after calling ChangeSlice()", prices)

	myMap := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap before calling ChangeMap():", myMap)
	ChangeMap(myMap)
	fmt.Println("myMap after calling ChangeMap():", myMap)

	myMap2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("myMap before calling ChangeMapByPointer():", myMap2)
	ChangeMapByPointer(&myMap2)
	fmt.Println("myMap after calling ChangeMapByPointer():", myMap2)

}
