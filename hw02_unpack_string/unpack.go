package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var builder strings.Builder
	switch {
	case len(s) > 0 && !unicode.IsDigit(rune(s[0])):
		for _, char := range s {
			if unicode.IsDigit(char) {
				repeat, _ := strconv.Atoi(string(char))
				str := builder.String()
				if repeat == 0 {
					builder.Reset()
					builder.WriteString(str[:len(str)-1])
				} else {
					builder.WriteString(strings.Repeat(str[len(str)-1:], repeat-1))
				}
			} else {
				builder.WriteString(string(char))
			}
		}
		return builder.String(), nil
	case len(s) > 0 && unicode.IsDigit(rune(s[0])):
		return "", ErrInvalidString
	default:
		return "", nil
	}
}
