package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// using the os package
	byteSlice := []byte("I am Ronald Benjamin and I learn GoLang")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		panic(err)
	}
	log.Printf("Byte Written %d", byteWritten)
	// os package till here

	// another way using go ioutil package
	bs := []byte("GO Programming is cool")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		panic(err)
	}
	// ioutil package
}
