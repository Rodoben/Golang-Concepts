package main

import "fmt"

func main() {

	var emp map[string]string

	fmt.Printf("%#v\n", emp)
	fmt.Printf("No of pairs %d\n", len(emp))

	//
	fmt.Printf("The value of key %q is %q\n", "dan", emp["dan"])
	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["abc"])

	//var mp map[[]int]string
	mp := map[[5]int]string{
		//    mp[[1,2,3,4]]:"abc",
	}
	_ = mp

	//	emp["dan"] = "Programmer"

	people := map[string]float64{}
	fmt.Println(people)

	people["john"] = 21.4
	people["jep"] = 2.4
	fmt.Println(people)

	maps1 := make(map[int]int)
	maps1[4] = 4
	fmt.Println(maps1)

	bal := map[string]int{
		"ron": 12,
		"ben": 34,
		"com": 90,
	}
	fmt.Println(bal)
	m := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println(m)

	bal["ggg"] = 55
	bal["ttt"] = 44
	fmt.Println(bal)
	fmt.Println(bal["Ron"])

	v, ok := bal["ron"]
	if ok {
		bal["ron"] = bal["ron"] + 1
		fmt.Println(bal["ron"])
	} else {
		fmt.Println(v)
	}

	for k, v := range bal {
		fmt.Printf("Key: %#v, val:%#v\n", k, v)
	}

	a := map[string]int{"a": 1}
	b := map[string]int{"a": 2}
	//fmt.Println(a == b)

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are Equal")

	} else {
		fmt.Println("Maps are unequal")
	}

	/////////////////////////////////////////
	// MAP HEADER{}

	friends := map[string]int{"Dan": 40, "Maria": 20}
	neighbours := friends
	friends["Dan"] = 1
	fmt.Println("Friends:", friends, "Neighbour", neighbours)

	// cloning map
	people1 := make(map[string]int)

	for k, v := range friends {
		people1[k] = v
	}
	fmt.Println("People1:", people1)
	people1["rodo"] = 34
	neighbours["ddd"] = 34
	fmt.Println("Friends:", friends, "Neighbour", neighbours, "People1", people1)

}
