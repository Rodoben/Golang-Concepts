package main

import "fmt"

func main() {

	// var m1 map[string]string
	// fmt.Printf("%#v -----%T\n", m1, m1)

	// m2 := map[int]string{1: "Ronald", 2: "Benjamin"}
	// _ = m2
	// m2[10] = "Abba"

	// fmt.Println(m2)

	// fmt.Println(m2[1])
	// fmt.Println(m2[100])

	/////////////////////////////////////////////
	var m1 map[int]bool
	//m1[5] = true
	_ = m1
	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	a := fmt.Sprintf("%s", m2)
	b := fmt.Sprintf("%s", m3)

	if a == b {
		fmt.Println("Maps are Equal")
	} else {
		fmt.Println("Maps are Unequal")
	}

	m := map[int]bool{1: true, 2: false, 3: false}

	delete(m, 1)
	for i, r := range m {
		fmt.Println(i, r)
	}

}
