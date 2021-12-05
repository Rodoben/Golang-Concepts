package main

import "fmt"

func main() {

	colors := make(map[int]string)

	colors[0] = "pink"
	colors[1] = "white"
	colors[2] = "red"
	colors[3] = "yellow"
	colors[4] = "green"

	f := "yellow"

	fmt.Println(check(colors, f))

}

func check(color map[int]string, f string) bool {
	for _, c := range color {

		if f == c {
			fmt.Print(f)
			fmt.Print(c)
			return true
		}
	}
	return false
}
