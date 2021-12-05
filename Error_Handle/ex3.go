// The panic keyword is used to terminate the program with a custom error message.
//When the panic keyword is encountered, the following set of instructions are followed:
// Execution for the current function stops
// Any function called with defer keyword is executed
// The program execution is terminated

package main

import "fmt"

func executePanic() {
	panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}
