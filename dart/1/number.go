package main

import (
	"fmt"
	"math/rand"
	"time"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {
	var k int

	myrand := random(0, 20)

	for k = 0; k < 3; k++ {

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Send your number: ")
		text, err := reader.ReadString('\n')
			if err != nil{
				fmt.Println(err)
			}

		text = strings.Trim(text, "\n")

		i, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println(err)
			}

		if myrand == i {
			fmt.Println("Number true because", i == myrand)
			break
		}else{
			if i > myrand {
				fmt.Println("Number less")
			}else{
				fmt.Println("Number more")
			}
		}
	}
	if k == 3 {
		fmt.Println("Your game over, number was:", myrand)
	}
}