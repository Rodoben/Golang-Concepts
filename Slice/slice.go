package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var cities []string
	fmt.Println("cities is equal to nil:", cities == nil)

	fmt.Printf("cities%#v \n", cities)
	fmt.Println(len(cities))
	//cities[0] = "London"
	////////////////////////////////////////////////////////////

	nums := []int{2, 9, 8, 9, 0}
	fmt.Println(nums)
	number := make([]int, 2)
	fmt.Printf("%#v\n", number)
	/////////////////////////////////////////////////////
	type name []string
	friends := name{"Ronald", "Ashutosh"}
	fmt.Println(friends)
	myFrnd := friends[0]
	fmt.Println("myFriend", myFrnd)
	fmt.Println("myFriend", friends[1])
	/////////////////////////////////////////////////
	//iteration over a slice
	for i, v := range nums {
		fmt.Printf("index: %v\n value: %v\n", i, v)
	}
	////////////////////////////////////////////////
	var n []int
	n = nums //copying a slice
	fmt.Println(n)

	//////////////////////////////////////////////////////
	//comparing a slice
	var n1 []int
	fmt.Println(n1 == nil)
	m := []int{}
	fmt.Println(m == nil)
	fmt.Println(m)

	a, b := []int{1, 2, 3}, []int{1, 8, 3}
	//fmt.Println(a == b) // slice can only be compared to nil

	// to compare slice

	var eq bool = true
	a = nil
	for i, valA := range a {
		if valA != b[i] {
			eq = false
			break
		}
	}
	if len(a) != len(b) {
		eq = false
	}
	if eq {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are unequal")
	}
	////////////////////////////////////////////////////////////////////

	n2 := []int{}
	n2 = append(n2, 10)
	fmt.Println(n2)
	n2 = append(n2, 20, 30, 40)
	fmt.Println(n2)
	///////////////////////////////////////////////////////////
	//when we want to append a slice to another slice
	n3 := []int{100, 200, 300}
	n2 = append(n2, n3...)
	fmt.Println(n2)
	/////////////////////////////////////////////////////
	// to copy  a slice to another slice
	src := []int{10, 20, 30, 40}
	dst1 := []int{50, 56}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn, dst1)
	//////////////////////////////////////////////////////////////////
	//Slice Expression:

	a1 := [5]int{1, 2, 3, 4, 5}
	b1 := a1[1:3]
	fmt.Printf("%v,%T\n", b1, b1)
	b2 := a1[3:5]
	fmt.Println(b2)
	fmt.Println(a1[2:])
	fmt.Println(a1[:3])
	fmt.Println(a1[:])
	//////////////////////////////////////////////////////////////////
	s1 := []int{10, 20, 30, 50, 40, 80}
	s1 = append(s1[:3], 200)
	fmt.Println(s1)
	s1 = append(s1[:3], 500)
	fmt.Println(s1)
	s1 = append(s1[:4], 900)
	fmt.Println(s1)
	///////////////////////////////////////////////////////////////////
	// Backing array
	ss1 := []int{10, 20, 30, 40, 50}
	ss2, ss4 := ss1[0:2], ss1[1:3]
	fmt.Println(ss2, ss4)
	ss2[1] = 900 // Modifying the slice using s2
	fmt.Println(ss1)

	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	fmt.Println(arr1, slice1, slice2)
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice2)
	////////////////////////////////////////////////////////////////////////////////////////////
	// To copy Existing slice to new slice
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}
	newCars = append(newCars, cars[1:3]...)
	fmt.Println(newCars)
	///////////////////////////////////////////////////////////////////////////////////////////
	f1 := []int{10, 20, 30, 40, 50, 60}
	nS := f1[0:3]
	fmt.Println(len(nS), cap(nS))

	nS = f1[2:4]
	fmt.Println(len(nS), cap(nS))
	nS = f1[3:6]
	fmt.Println(len(nS), cap(nS))
	//////////////////////////////////////////////////////////////////////////////////////////

	ar11 := [5]int{1, 2, 3, 4, 5}
	sl11 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Array Size %d \n", unsafe.Sizeof(ar11))
	fmt.Printf("Slice Size %d \n", unsafe.Sizeof(sl11))
	///////////////////////////////////////////////////////////
	//Append, length and capacity in depth

	var num1 []int
	fmt.Printf("%#v \n", num1)
	fmt.Printf("Length: %d,Capacity: %d\n", len(num1), cap(num1))
	num1 = append(num1, 1, 2, 3)
	fmt.Printf("Length: %d,Capacity: %d\n", len(num1), cap(num1))
	num1 = append(num1, 5)
	fmt.Printf("Length: %d,Capacity: %d\n", len(num1), cap(num1))
	num1 = append(num1, 4, 6)
	fmt.Printf("Length: %d,Capacity: %d\n", len(num1), cap(num1))
	num1 = append(num1, 56)
	fmt.Printf("Length: %d,Capacity: %d\n", len(num1), cap(num1))
	/////////////////////////////////////////////////////////////////////////

	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y")
	fmt.Println(letters, len(letters), cap(letters))
	fmt.Println(letters[3:6], len(letters), cap(letters))
	letters = append(letters, "R", "B")
	fmt.Println(letters[3:6], len(letters), cap(letters))

	//sfmt.Println(letters[4])

}
