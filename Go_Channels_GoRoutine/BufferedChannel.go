package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Println("func goroutine starts(before) sending data into ther channel", i)
			c <- i
			fmt.Println("func gorutine after sending data into the channel", i)
		}
		close(c)
	}(c)

	fmt.Println("Main Goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c {
		fmt.Println("main goroutine receive value from channel", v)
	}

}
