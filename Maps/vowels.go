package main

import (
	"fmt"
)

func main() {

	d := []string{"o", "u", "u", "r", "a", "n", "g", "e", "e"}

	vowel := map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}

	for _, r := range d {

		_, ok := vowel[r]
		if ok {
			vowel[r]++
			//fmt.Println(r, vowel[r])
		} else {

		}

	}

	for i, r1 := range vowel {

		if r1 != 0 {
			fmt.Println(i, r1)
		}

	}

}
