package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const g = 100
	var wg sync.WaitGroup
	wg.Add(g * 2)
	var n int = 0
	for i := 0; i < g; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)

}
