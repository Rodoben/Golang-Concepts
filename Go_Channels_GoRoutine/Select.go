package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "Hello"

	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Ronald"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Receieved:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)

			//default:
			//fmt.Println("no activity")
		}
	}

}
