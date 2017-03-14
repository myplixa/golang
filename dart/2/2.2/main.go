package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"log"
)

func main() {

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}
