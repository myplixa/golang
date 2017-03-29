package calculate

import (
	"io"
	"bufio"
)

func Calc(stream io.Reader) int {

	reader := bufio.NewReader(stream)
	count := 0;

	for {
		_, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		count ++
	}
	return count
}