package main

import (
	"fmt"
	"io/ioutil"
)

func readFiles() string {

	read, err := ioutil.ReadFile("./text/files_1.txt")
	if err != nil {
		panic(err)
	}
	return(string(read))
}

func main()  {

	letters := []string{".", ",", "(", ")", "'", ":", ";", "-", "!", "?", "*", "\n", "\t", "\r"}

	fmt.Print(letters)

	fmt.Print(readFiles())

}