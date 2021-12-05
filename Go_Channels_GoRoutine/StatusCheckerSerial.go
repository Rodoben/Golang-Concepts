package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://manikchugh.xyz",
	}

	for _, link := range links {
		CheckLinks(link)
	}

}

func CheckLinks(link string) {

	a, err := http.Get(link)
	if err != nil {

		fmt.Println(link, " might be down", err)
		return
	}
	fmt.Println(link, " is up", a)

}
