package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"time"
	"io/ioutil"
	"strings"
	"math/rand"
)



func enterLetter() string {
	reader := bufio.NewReader(os.Stdin)

	l, e := reader.ReadByte()
	if e != nil {
		log.Fatalf("ERR(reading letter from user): %s", e.Error())
	}

	return string(l)
}

func findLetter(r, word string) bool {

	for _, f := range word {

		if (string(f) == r) {
			return true
		}

	}
	return false
}

func iswinner(enttrue, word string) bool {

	for _, f := range word {

		if !findLetter(string(f), enttrue) {
			return false
		}
	}
	return true
}

func wordstring(enttrue, word string) (ak string) {

	for _, f := range word {

		if (findLetter(string(f), enttrue)) {
			ak = ak + string(f)
		} else {
			ak = ak + "_"
		}
	}
	return
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func randomWord() string {

	r, err := ioutil.ReadFile("./hangame.txt")
	if err != nil {
		fmt.Println(err)
	}

	w := strings.Split(string(r), "\n")
	myrand := random(0, (len(w) - 1))

	return w[myrand]

}


func main() {

	var k int

	word := randomWord()

	var entTru string


	for k = 0; k < len(word)+2; k ++ {

		fmt.Print("введите символ: ")
		text := enterLetter()

		q := findLetter(text, word)

		if q {
			fmt.Println("симовл есть")
			entTru = entTru + text
		} else {
			fmt.Println("символа нет")
		}
		fmt.Printf("status %s \n", wordstring(entTru, word))

		if iswinner(entTru, word) {
			break
		}
	}

	if iswinner(entTru, word) {
		fmt.Println("You winner")
	} else {
		fmt.Println("You LoL")
	}
}