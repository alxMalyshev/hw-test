package hw09structvalidator

import (
	"errors"
)

var (
	ErrStringLengthInvalid = errors.New("string is not equal len rule")
	ErrIntMinInvalid       = errors.New("int less then it allows in min rule")
	ErrIntMaxInvalid       = errors.New("int more then it allows in max rule")
	ErrStringRegexpInvalid = errors.New("string is not compatible with regexp rule")
	ErrStringInInvalid     = errors.New("string is not compatible with \"in\" rule")
	ErrIntInInvalid        = errors.New("int is not compatible with \"in\" rule")

	ErrFieldTypeInvalid = errors.New("filed value has invalid type")

	ErrStringRegexpFormat = errors.New("string \"regexp\" tag rule format icorrect")
	ErrStringLengthFormat = errors.New("string \"len\" tag rule format incorrect")
	ErrIntMinFormat       = errors.New("int \"min\" tag rule format incorrect")
	ErrIntMaxFormat       = errors.New("int \"max\" tag rule format incorrect")
	ErrStringInFormat     = errors.New("string \"in\" tag rule format incorrect")
	ErrIntInFormat        = errors.New("int \"in\" tag rule format incorrect")

	ErrStructValidation = errors.New("non struct type pass for validation")
	ErrKindValidation   = errors.New("incompatible kind of value filed")
)
