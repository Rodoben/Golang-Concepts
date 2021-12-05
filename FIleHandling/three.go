package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	// Writing using bufio package
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{97, 98, 99}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written to buffer(not File", bytesWritten)

	log.Printf("Bytes Written to buffer(not File", bufferedWriter.Available())
}
