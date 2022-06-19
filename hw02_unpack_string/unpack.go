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
	val := []rune(s)
	switch {
	case len(val) > 0 && !unicode.IsDigit(val[0]):
		for i := 0; i <= len(val)-1; i++ {
			switch {
			case unicode.IsDigit(val[i]):
				if lastState == "char" {
					repeat, _ := strconv.Atoi(string(val[i]))
					if repeat == 0 {
						str := []rune(builder.String())
						builder.Reset()
						builder.WriteString(string(str[:len(str)-1]))
					} else {
						builder.WriteString(strings.Repeat(string(val[i-1]), repeat-1))
					}
				} else if lastState == "backslash" {
					builder.WriteString(string(val[i]))
					lastState = "char"
					continue
				} else {
					return "", ErrInvalidString
				}
				lastState = "num"
			case string(val[i]) == `\`:
				if lastState == "backslash" {
					builder.WriteString(string(val[i]))
					lastState = "char"
				} else {
					lastState = "backslash"
					continue
				}
			default:
				builder.WriteString(string(val[i]))
				lastState = "char"
			}
		}
		return builder.String(), nil
	case len(val) > 0 && unicode.IsDigit(val[0]):
		return "", ErrInvalidString
	default:
		return "", nil
	}
}
