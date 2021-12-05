package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")

	if err != nil {
		fmt.Println("Error", err)
	}
	log.SetOutput(nf)
}

func main() {
	nf, err := os.Open("no-file1.txt")
	if err != nil {
		//fmt.Println("error happend", err)
		//log.Println("err happnd", err)
		//log.Fatal(err)
		panic(err)

	}
	fmt.Println(nf, err)
}
