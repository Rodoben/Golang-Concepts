package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{

		"http://google.com",
		"http://facebook.com",
		"http://amazon.in",
		"http://manikchugh.xyz",
		"http://flipkart.com",
		"http://ronald-benjamin.web.app",
	}
	c := make(chan string, 3)

	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		fmt.Println("*****************")
		go checkLink(<-c, c)
	}

}

func checkLink(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		c <- link

		fmt.Println(link, "might be down")
		return

	}
	c <- link

	fmt.Println(link, "is up")

}
