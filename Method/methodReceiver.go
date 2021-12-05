package main

import "fmt"

type contact struct {
	email, addr string
	phno        int
}

type Employee struct {
	name  string
	age   int
	cInfo contact
}

type EmpList []Employee

func CreateRecord() EmpList {
	var name, email, addr string
	var phno, age int
	nb := []Employee{}
	ch := 0
	var c int
	for ch == 0 {
		fmt.Println("Enter Employee Name")
		fmt.Scanln(&name)
		fmt.Println("Enter Phone Number")
		fmt.Scanln(&phno)
		fmt.Println("Enter Employee addr")
		fmt.Scanln(&addr)
		fmt.Println("Enter Employee email")
		fmt.Scanln(&email)
		fmt.Println("Enter Employee age")
		fmt.Scanln(&age)

		b := Employee{

			name: name,
			age:  age,
			cInfo: contact{
				email: email,
				addr:  addr,
				phno:  phno,
			},
		}
		nb = append(nb, b)
		fmt.Println("Do you want to enter more records? \n 0.Yes \n 1.No")
		fmt.Scanln(&c)
		ch = c

	}
	return nb
}

func (e EmpList) print() {

	for i, list := range e {
		fmt.Println(i, list)

	}
	fmt.Printf("Type: %T \n", e)
}

func (e EmpList) StackOp() {
	var ch int
	fmt.Println("Welcome to stack")
	fmt.Println("1. push \n 2.Pop \n 3.Display")
	fmt.Scanln(&ch)
	switch ch {
	case 1:
		CreateRecord()

	case 2:

	case 3:
		CreateRecord().print()
	default:

		fmt.Println("wrong choice")
	}

}

func main() {
	EmpList := CreateRecord()
	EmpList.print()
	//EmpList.StackOp()

}
