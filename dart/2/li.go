package main


import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
)


func readLetter() (string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter letter: ")
	l, e := reader.ReadByte()

	if e != nil {
		log.Fatalf("ERR(reading letter from user): %s", e.Error())
	}

	return string(l)
}

func main() {

	// заданное слово.
	v := "book"
	// масивы для ввода данных.
	var entTru []string
	var entFols []string

	for _,f := range v {

		if(string(f) == v){

			fmt.Print(string(f))
			entTru = append(entTru, string(v))

		} else {

			fmt.Print("_")
		}

		//струкура для ввода символов
		text := readLetter()

		r := strings.Contains(v, string(text))

		if r {

			fmt.Println("совпадений найдено =", strings.Count(v, string(text)))
			entTru = append(entTru, string(text))
			fmt.Println(entTru)

		} else {

			entFols = append(entFols, string(text))
			fmt.Println(entFols)
			fmt.Println("совпадений не найдено, попробуйте снова.")
		}
	}
}