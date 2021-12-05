package main

import (
	"fmt"
)

func main() {
	str := "Ronald"
	str1 := "is"

	runes := []rune(str)
	fmt.Println(string(runes[:3]) + str1 + string(runes[3:]))

}
