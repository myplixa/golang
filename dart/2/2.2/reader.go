package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFiles() string {

	read, err := ioutil.ReadFile("./text/files_1.txt")
	if err != nil {
		panic(err)
	}
	return(string(read))
}

//func EditFiles(r, letters) string  {
//
//}

func main()  {

	letters := []string{".", ",", "(", ")", "'", ":", ";", "-", "–", "!", "?", "*", "\n", "\t", "\r"}

	lower := strings.ToLower(readFiles())

	for f := range letters {
		lower = strings.Replace(lower, letters[f], " ", -1)
	}

	fmt.Println(lower)

}