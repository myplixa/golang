package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	leters := []string{",", ".", "(", ")", ";", ":", "-", "/", "\\", "\"", "'", "*", "?", "!", "\t", "\r", "\n"}

	//причитатит файл
	bs, err := ioutil.ReadFile("./text.txt")
	if err != nil {
		return
	}

	//привести данные к нижнему регистпу
	bss := strings.ToLower(string(bs))
	//заменяем все спец символы на пробел используя цыкл
	for k := range leters {
		bss = strings.Replace(bss, leters[k], " ", -1)

	}
	//переводим строки в масив с разделением по пробелу
	arr := strings.Split(bss, " ");
	//создпние map для выводы слов больше 3-х символов
	bsx := make(map[string]int)
	for _, v := range arr {
		l := strings.Split(v, "");
		if(len(l) >=3){
			_, ok := bsx[v]
			if ok {
				bsx[v] = bsx[v] +1
			} else {
				bsx[v] = 1
			}
		}
	}
	fmt.Println(bsx)
}