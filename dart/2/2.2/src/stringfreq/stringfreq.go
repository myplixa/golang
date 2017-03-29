package stringfreq

import (
	"io"
	"bytes"
)


func StringFreq(stream io.Reader) string {

	reader := new(bytes.Buffer)
	reader.ReadFrom(stream)

	return reader.String()
}