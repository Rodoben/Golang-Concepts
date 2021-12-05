package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "one two three"
	StringReverse(str)
}

func StringReverse(s string) {
	var arr []string

	words := strings.Fields(s)

	for _, r := range words {
		arr = append(arr, r)
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i], " ")
	}

}
