package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


func main() {

	var p int

	t := "door"

	var entTru []string
//	var entFols []string

	for p = 0; p < 9; p++ {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter letter: ")
		text, err := reader.ReadString('\n')

			if err != nil {

				fmt.Println(err)

			}

		text = strings.Trim(text, "\n")
		fmt.Print("[")

		for _, f := range t {

			if (string(f) == text) {
				fmt.Print(string(f))
				entTru = append(entTru, string(text))
			} else {
				fmt.Print("_")
			}
		}

		fmt.Print("]", "\n")

		fmt.Println(entTru)

	}
	if p == 9 {
		fmt.Println("Your game over, guesser loses - the answer was -", t)
	}
}
