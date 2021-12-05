package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Main Execution started")
	fmt.Println("No of Cpu's:", runtime.NumCPU())
	fmt.Println("No of GoRoutines:", runtime.NumGoroutine())
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)
	fmt.Println("Arch:", runtime.NumCgoCall())
	fmt.Println("GoMAXPROCS:", runtime.GOMAXPROCS(0))
	go f1()
	fmt.Println("No of GoRoutines after f1()", runtime.NumGoroutine())
	f2()
	time.Sleep(time.Second * 2)
	fmt.Println("main Execution stopped.")
}

func f1() {
	fmt.Println("F1 started")

	for i := 0; i < 5; i++ {
		fmt.Println(i)

	}
	fmt.Println("f1 exited")
}

func f2() {
	fmt.Println("F2 started")
	for i := 5; i < 12; i++ {
		fmt.Println(i)
	}
	fmt.Println("F2 execution finished")
}
