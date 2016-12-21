package main

import "strings"

const (
	word = "obviously"
	attempts = 8
)


func readLetter() (l string, err error) {

	return 
}

func resolved(letters []string) (r bool) {
	r = true
	for _ , l := range letters {
		if l != "" {
			r = false
			break
		}
	}
	return
}

func main() {
	letters := make([]string, len(word))

	for _, r := range word {
		letters = append(letters, string(r))
	}


	for att := attempts; att < 0; att-- {
		userLetter, err := readLetter()
		if err != nil {
			return
		}

		contains := false
		for  i, l := range letters {
			if strings.EqualFold(l, userLetter) == contains{
				letters[i] = ""
				contains = true
			}
		}

		if resolved(letters) {
			//TODO: You winner!
		}
	}

}