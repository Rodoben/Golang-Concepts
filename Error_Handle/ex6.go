package main

import (
	"fmt"
	"log"
	"os"
)

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println("Recovery Message:", recoveryMessage)
		nf, err := os.Create("no-file1.txt")

		if err != nil {
			fmt.Println("Error", err)
		}
		log.SetOutput(nf)

	}
	fmt.Println("This is recovery function...")
}

func executePanic() {
	defer recoveryFunction()
	nf, err := os.Open("no-file1.txt")
	if err != nil {
		panic(err)

	}

	fmt.Println("The function executes Completely", nf, err)
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}
