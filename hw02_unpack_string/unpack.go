package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {

	if s != "" && !unicode.IsDigit(rune(s[0])) {
		var str string

		for _, char := range s {

			if unicode.IsDigit(char) {
				repeart, _ := strconv.Atoi(string(char))
				if repeart == 0 {
					str = str[:len(str)-1]
				} else {
					str = str + strings.Repeat(str[len(str)-1:], repeart-1)
				}
			} else {
				str += string(char)
			}
		}
		return str, nil
	} else if s != "" && unicode.IsDigit(rune(s[0])) {
		return "", ErrInvalidString
	}
	return "", nil
}
