package hw09structvalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rules struct {
	Length interface{}
	Max    interface{}
	Min    interface{}
	Regexp interface{}
	In     interface{}
}

func (r *Rules) RuleStringLength(fieldValue reflect.Value) error {
	length, err := strconv.Atoi(r.Length.(string))
	if err != nil {
		return fmt.Errorf("%w:\n\tCause by: rule has wrong format %s", ErrStringLengthFormat, r.Length)
	}

	if len(fieldValue.String()) != length {
		return fmt.Errorf("%w:\n\tCause by: value %s not equals %s", ErrStringLengthInvalid, fieldValue.String(), r.Length)
	}

	return nil
}

func (r *Rules) RuleStringRegexp(fieldValue reflect.Value) error {
	reg, err := regexp.Compile(r.Regexp.(string))
	if err != nil {
		return fmt.Errorf("%w:\n\tCause by: rule wrong has format", ErrStringRegexpFormat)
	}

	if !reg.MatchString(fieldValue.String()) {
		return fmt.Errorf("%w:\n\tCause by: incompatible string with regexp %s", ErrStringRegexpInvalid, r.Regexp)
	}

	return nil
}

func (r *Rules) RuleIntMax(fieldValue reflect.Value) error {
	max, err := strconv.Atoi(r.Max.(string))
	if err != nil {
		return fmt.Errorf("%w:\n\tCause by: rule has wrong format %s", ErrIntMaxFormat, r.Max)
	}

	if int(fieldValue.Int()) > max {
		return fmt.Errorf("%w:\n\tCause by: field more then %s", ErrIntMaxInvalid, r.Max)
	}
	return nil
}

func (r *Rules) RuleIntMin(fieldValue reflect.Value) error {
	min, err := strconv.Atoi(r.Min.(string))
	if err != nil {
		return fmt.Errorf("%w:\n\tCause by: rule wrong format %s", ErrIntMinFormat, r.Min)
	}
	if int(fieldValue.Int()) < min {
		return fmt.Errorf("%w:\n\tCause by: field less then %s", ErrIntMinInvalid, r.Min)
	}
	return nil
}

func (r *Rules) RuleIntIn(fieldValue reflect.Value) error {
	ruleList := strings.Split(r.In.(string), ",")

	if len(ruleList) == 0 {
		return fmt.Errorf("%w\n\tCause by: rule has wrong format %s", ErrStringInFormat, r.In)
	}

	var intRule int
	var err error
	for _, rule := range ruleList {
		if intRule, err = strconv.Atoi(rule); err != nil {
			return err
		}

		if intRule == int(fieldValue.Int()) {
			return nil
		}
	}

	return fmt.Errorf("%w\n\tCause by: field value is not equal item of slice %s", ErrStringInInvalid, ruleList)
}

func (r *Rules) RuleStringIn(fieldValue reflect.Value) error {
	ruleList := strings.Split(r.In.(string), ",")
	valueList := strings.Split(fieldValue.String(), ",")

	if len(ruleList) == 0 {
		return fmt.Errorf("%w\n\tCause by: rule has wrong format %s", ErrStringInFormat, r.In)
	}

	for _, rule := range ruleList {
		for _, value := range valueList {
			if rule == value {
				return nil
			}
		}
	}

	return fmt.Errorf("%w\n\tCause by: field value is not equal item of slice %s", ErrStringInInvalid, ruleList)
}

func ParseRule(rule string) string {
	ruleValue := strings.Split(rule, ":")
	return ruleValue[1]
}
