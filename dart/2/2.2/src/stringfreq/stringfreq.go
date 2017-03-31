package stringfreq

import (
	"io"
	"bufio"
)

func Calc(stream io.Reader) map[string]int {

	reader := bufio.NewReader(stream)
	stringfreg := make(map[string]int)

	for  {
		word, err := ReadWord(reader)

		if i, ok := stringfreg[word]; ok {
			i ++
			stringfreg[word] = i
		} else {
			stringfreg[word] = 1
		}
	if err != nil {
		break
		}
	}
	return stringfreg
}

func ReadWord(reader *bufio.Reader) (string, error) {

}