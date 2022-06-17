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
	var lastState string
	switch {
	case len(s) > 0 && !unicode.IsDigit(rune(s[0])):
		for _, char := range s {
			switch {
			case unicode.IsDigit(char):
				str := builder.String()
				if lastState == "char" {
					repeat, _ := strconv.Atoi(string(char))
					if repeat == 0 {
						builder.Reset()
						builder.WriteString(str[:len(str)-1])
					} else {
						builder.WriteString(strings.Repeat(str[len(str)-1:], repeat-1))
					}
				} else if lastState == "backslash" {
					builder.WriteString(string(char))
					lastState = "char"
					continue
				} else {
					return "", ErrInvalidString
				}
				lastState = "num"
			case string(char) == `\`:
				if lastState == "backslash" {
					builder.WriteString(string(char))
					lastState = "char"
				} else {
					lastState = "backslash"
					continue
				}
			default:
				builder.WriteString(string(char))
				lastState = "char"
			}
		}
		return builder.String(), nil
	case len(s) > 0 && unicode.IsDigit(rune(s[0])):
		return "", ErrInvalidString
	default:
		return "", nil
	}
}
