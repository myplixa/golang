package main

import (
	"os"
//	"calculate"
	"fmt"
//	"runefreq"
	"stringfreq"
	"calculate"
	"runefreq"
)

func main() {

	i := calculate.Calc(os.Stdin)
	fmt.Printf("the number of characters = %d\n",i)

 	m := runefreq.Calc(os.Stdin)
	for r, c := range m {
		fmt.Printf("%#U = %d\n", r, c)
	}

 	s := stringfreq.StringFreq(os.Stdin)
	fmt.Println(s)
}
