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
	const (
		statusChar      = "char"
		statusBackSlash = "backslash"
		statusNumber    = "num"
	)
	switch {
	case len(val) > 0 && !unicode.IsDigit(val[0]):
		for i := 0; i <= len(val)-1; i++ {
			switch {
			case unicode.IsDigit(val[i]):
				switch {
				case lastState == statusChar:
					repeat, err := strconv.Atoi(string(val[i]))
					if err != nil {
						return errors.New("Converting Error")
					}
					if repeat == 0 {
						str := []rune(builder.String())
						builder.Reset()
						builder.WriteString(string(str[:len(str)-1]))
					} else {
						builder.WriteString(strings.Repeat(string(val[i-1]), repeat-1))
					}
				case lastState == statusBackSlash:
					builder.WriteString(string(val[i]))
					lastState = statusChar
					continue
				default:
					return "", ErrInvalidString
				}
				lastState = statusNumber
			case string(val[i]) == `\`:
				if lastState == statusBackSlash {
					builder.WriteString(string(val[i]))
					lastState = statusChar
				} else {
					lastState = statusBackSlash
					continue
				}
			default:
				builder.WriteString(string(val[i]))
				lastState = statusChar
			}
		}
		return builder.String(), nil
	case len(val) > 0 && unicode.IsDigit(val[0]):
		return "", ErrInvalidString
	default:
		return "", nil
	}
}
