package hw09structvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "validate"

type ValidationError struct {
	Field string
	Err   error
}

type Field struct {
	Name  string
	Value reflect.Value
	Tag   string
	Type  reflect.StructField
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var builder strings.Builder

	builder.WriteString("Stack: \n")
	for _, err := range v {
		builder.WriteString("Field ")
		builder.WriteString(err.Field)
		builder.WriteString(":\n\t")
		builder.WriteString(err.Err.Error())
		builder.WriteString("\n")
	}
	return builder.String()
}

func (f *Field) processFieldByRule(ruleList []string) *ValidationError {
	rules := Rules{}
	var wrapErr error

	for _, rule := range ruleList {
		switch {
		case strings.HasPrefix(rule, "len"):
			rules.Length = ParseRule(rule)
			if err := rules.RuleStringLength(f.Value); err != nil {
				wrapErr = fmt.Errorf("%w", err)
			}
		case strings.HasPrefix(rule, "min"):
			rules.Min = ParseRule(rule)
			if err := rules.RuleIntMin(f.Value); err != nil {
				wrapErr = fmt.Errorf("%w", err)
			}
		case strings.HasPrefix(rule, "max"):
			rules.Max = ParseRule(rule)
			if err := rules.RuleIntMax(f.Value); err != nil {
				wrapErr = fmt.Errorf("%w", err)
			}
		case strings.HasPrefix(rule, "regexp"):
			rules.Regexp = ParseRule(rule)
			if err := rules.RuleStringRegexp(f.Value); err != nil {
				wrapErr = fmt.Errorf("%w", err)
			}
		case strings.HasPrefix(rule, "in"):
			rules.In = ParseRule(rule)
			var errIn error

			if f.Value.Kind() == reflect.Int {
				errIn = rules.RuleIntIn(f.Value)
			}
			if f.Value.Kind() == reflect.String {
				errIn = rules.RuleStringIn(f.Value)
			}

			if errIn != nil {
				wrapErr = fmt.Errorf("%w", errIn)
			}
		}
	}

	return &ValidationError{
		Field: f.Name,
		Err:   wrapErr,
	}
}

func Validate(v interface{}) error {
	var errs ValidationErrors
	obj := reflect.ValueOf(v)
	field := new(Field)

	if obj.Kind() != reflect.Struct {
		return fmt.Errorf("%w\n\t Cause by: expected struct, but received %T", ErrStructValidation, v)
	}

	var value reflect.Value
	for i := 0; i <= obj.NumField()-1; i++ {
		field.Name = obj.Type().Field(i).Name
		value = obj.Field(i)
		field.Tag = obj.Type().Field(i).Tag.Get(tagName)
		field.Type = obj.Type().Field(i)

		if field.Tag == "" {
			continue
		}

		rules := strings.Split(field.Tag, "|")
		if len(rules) == 0 {
			return fmt.Errorf("\n\t Cause by: rule is empty")
		}

		//nolint
		switch value.Kind() {
		case reflect.Array, reflect.Slice:
			for i := 0; i < value.Len(); i++ {
				field.Value = value.Index(i)
				if err := field.processFieldByRule(rules); err.Err != nil {
					errs = append(errs, *err)
				}
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.String:
			field.Value = value
			if err := field.processFieldByRule(rules); err.Err != nil {
				errs = append(errs, *err)
			}
		default:
			errs = append(errs, ValidationError{
				field.Name,
				fmt.Errorf("%w: field is not int, string or array", ErrKindValidation),
			})
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}
