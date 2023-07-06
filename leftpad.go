package leftpad

import "bytes"

func Leftpad(str string, length int, ch byte) string {
	if len(str) >= length {
		return str
	}

	return string(bytes.Repeat([]byte{ch}, length-len(str))) + str
}
