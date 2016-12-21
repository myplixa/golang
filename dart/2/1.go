package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

func readLetter() (string) {
	reader := bufio.NewReader(os.Stdin)

	l, e := reader.ReadByte()
	if e != nil {
		log.Fatalf("ERR(reading letter from user): %s", e.Error())
	}

	return string(l)
}

func main()  {

	t := "door"

	var entTru []string

	fmt.Print("Enter letter: ")
	text := readLetter()

	for _, f := range t {

		if(string(f) == text) {
			fmt.Printf("[ %s ] \n", string(f))
			entTru = append(entTru, text)
		} else {
			fmt.Print("_")
		}
	}
}
