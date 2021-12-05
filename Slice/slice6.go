package main

import "fmt"

func main() {
	sum := 0
	nums := []int{5, -1, 9, 10, 110, 6, -1, 6}
	for _, r := range nums[2 : len(nums)-2] {
		sum = sum + r
		fmt.Println(r)
	}
	fmt.Println(sum)
}
