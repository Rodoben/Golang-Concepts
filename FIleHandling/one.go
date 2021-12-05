package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T \n", newFile)

	newFile, err := os.Create("as")
	if err != nil {
		panic(err)
	}

	err = os.Truncate("a.txt", 2)

	file, err := os.OpenFile("a.txt", os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	defer file.Close()
	fmt.Println(file)

	// check the status of the file

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
	p("FileName:", fileInfo.Name())
	p("size in Bytes", fileInfo.Size())
	p("Last Modified:", fileInfo.ModTime())
	p("Is Directory?", fileInfo.IsDir())
	p("Permission:", fileInfo.Mode())
	p(fileInfo.Sys())



	fileInfo, err = os.Stat("a.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File Does not exist")
		} else {
			log.Fatal("error Exsits")
		}
	}

	// Rename a file

	oldpath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldpath, newPath)
	if err != nil {
		log.Fatal(err)
	}


	// Remove a file
	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

}
