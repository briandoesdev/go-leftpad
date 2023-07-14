package leftpad

import "bytes"

func Leftpad(str string, length int, ch rune) string {
	if len(str) >= length {
		return str
	}

	// return leftpad string using rune instead of byte
	return string(bytes.Repeat([]byte(string(ch)), length-len(str))) + str
}
