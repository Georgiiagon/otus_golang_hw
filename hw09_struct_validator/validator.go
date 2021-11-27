package hw09structvalidator

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var result string
	for _, err := range v {
		result = result + " " + err.Field + ": " + err.Err.Error()
	}
	return result
}

var validationErrors = make(ValidationErrors, 0)

func Validate(v interface{}) error {
	reflectV := reflect.ValueOf(v)
	typeV := reflectV.Type()
	for i := 0; i < reflectV.NumField(); i++ {
		value := reflectV.Field(i)
		field := typeV.Field(i)
		tagV := field.Tag
		rules := SplitRules(tagV.Get("validate"))
		validateField(field.Name, value, rules)
	}
	fmt.Println(len(validationErrors), validationErrors)
	return nil
}

func validateField(fieldName string, value reflect.Value, rules []Rule) {
	switch value.Kind() { //nolint:exhaustive
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			validateField(fieldName, value.Index(i), rules)
		}
		return
	case reflect.String:
		for _, rule := range rules {
			err := rule.ValidateString(value)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
			}
		}
	case reflect.Int:
		for _, rule := range rules {
			err := rule.ValidateInt(value)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
			}
		}
	}
}
