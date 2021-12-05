package main

import "fmt"

func main() {
	nums := []float64{10.23, 52.32, 632.2}
	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{45.5, 65.7}
	nums = append(nums, n...)
	fmt.Println(nums)
}
