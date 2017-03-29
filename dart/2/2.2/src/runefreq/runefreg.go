package runefreq

import (
	"io"
	"bufio"
)

func Calc(stream io.Reader) map[rune]int {

	reader := bufio.NewReader(stream)
	freg := make(map[rune]int)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if i, ok := freg[r]; ok {
			i ++
			freg[r] = i
		} else {
			freg[r] = 1
		}
	}
	return freg
}
