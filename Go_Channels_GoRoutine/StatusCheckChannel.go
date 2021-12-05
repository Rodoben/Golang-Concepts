package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan string)
	links := []string{

		"http://google.com",
		"http://facebook.com",
		"http://ronald-benjamin.web.app",
	}

	for _, link := range links {
		go checkLink(link, c)
	}
	//fmt.Println(<-c)
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {

		fmt.Println(link, " is down")
		c <- "is up now"
	}
	fmt.Println(link, "is up")
	c <- "is up"
}
