package main

import (
	"fmt"
	"sync"
)

func main() {

	var m sync.Mutex

	const g = 100
	n := 0
	for i := 0; i < g; i++ {
		go func() {
			m.Lock()
			n++
			m.Unlock()
		}()

		go func() {
			m.Lock()
			n--
			m.Unlock()
		}()

	}

	fmt.Println(n)

}
