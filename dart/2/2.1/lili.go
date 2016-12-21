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

func main() {

	for {
		fmt.Print("введите символ:")
		text := readLetter()
		fmt.Printf("вы ввели = %s \n", text)
	}
}
