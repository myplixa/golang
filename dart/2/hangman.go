package main

import (
	"fmt"
	"io/ioutil"
	"time"
	"math/rand"
	"strings"
	"os"
	"bufio"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {

	var p int

	//читаем файл
	r, err := ioutil.ReadFile("./hangame.txt")
	if err != nil {
		fmt.Println(err)
	}

	w := strings.Split(string(r), "\n")
	myrand := random(0, (len(w) - 1))

	//тестовый вывод//
	fmt.Println("Outup from random:", w[myrand])
	//разбиваем текс по пустому символу
	t := strings.Split(w[myrand], "")

	//тестовый выыод
	fmt.Println("output from array:", t)

	for p = 0; p < 9; p++ {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter letter: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		text = strings.Trim(text, "\n")

	}
	if p == 9 {
		fmt.Println("Your game over, guesser loses - the answer was -", w[myrand])
	}
}