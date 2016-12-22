package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var entTru []string

func enterLetter() string {
	reader := bufio.NewReader(os.Stdin)

	l, e := reader.ReadByte()
	if e != nil {
		log.Fatalf("ERR(reading letter from user): %s", e.Error())
	}

	return string(l)
}

func findLetter() string {

	word := "door"

	for _, f := range word {

		if (string(f) == enterLetter()) {
			entTru = append(entTru, enterLetter())
		} else {
			entTru = append(entTru, "_")
		}
	}
	return findLetter()
}

func main() {

	for {
		fmt.Print("введите символ: ")
		text := enterLetter()
		fmt.Printf("вы ввели = %s \n", text)
		fmt.Println(entTru)
	}
}
