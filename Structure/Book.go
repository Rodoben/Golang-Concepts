package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// declarinng a new variable
type bookCredentials struct {
	title, author string
	year          int
}

//array of books
type book []bookCredentials

func newBook() book {
	nb := []bookCredentials{}
	var n int
	fmt.Print("enter number of entries: ")
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		var t, a string
		var y int
		fmt.Println("Enter Book Title")
		fmt.Scanln(&t)
		fmt.Println("Enter Book Author")
		fmt.Scanln(&a)
		fmt.Println("Enter Book year")
		fmt.Scanln(&y)

		b := bookCredentials{
			title:  t,
			author: a,
			year:   y,
		}
		nb = append(nb, b)
	}
	return nb
}

func (b book) print() {
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i].title, " ", b[i].author, " ", b[i].year)
	}
	// fmt.Println(b.title, " ", b.author, " ", b.year)
}

func (b book) toString() string {
	fmt.Println("to string conversion")
	var s []string
	for _, v := range b {

		s = append(s, v.title, v.author, strconv.Itoa(v.year))

		fmt.Println("Srtring Conversion", s)
	}
	justString := strings.Join(s, " ")
	fmt.Println("reached end point")
	return justString
}

func (b book) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(b.toString()), 0666)
}

func main() {

	var ch int
	fmt.Println("Welcome to Book shelf")
	for {
		fmt.Println("****MENU***** \n 1.Create a New Book Record \n 2.Display All Books \n 3.SaveTOFile")
		fmt.Scanln(&ch)
		switch ch {
		case 1:
			newBook()
		case 2:
			print()

		case 3:

		default:
			fmt.Println("Opps Wrong Choice")
		}
	}
}
