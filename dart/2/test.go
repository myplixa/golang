package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {

	n := "doc"

	var p int

	var entTru []string
	var entFols []string

	for p = 0; p < 9; p++ {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter letter: ")
		text, err := reader.ReadByte()

		if err != nil {

			fmt.Println(err)

		}

		i := strings.Contains(n, string(text))

		if i == true {

			entTru = append(entTru, string(text))
			fmt.Println(entTru)

		} else {

			entFols = append(entFols, string(text))
			fmt.Println(entFols)

		}
	fmt.Println(i)

	}
	if p == 9 {
		fmt.Println("Game Over, word was =", n)
	}
}