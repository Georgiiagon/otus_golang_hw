package hw09structvalidator

import (
	"errors"
	"log"
	"reflect"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var tagValue = "validate"

func (v ValidationError) Error() string {
	var b strings.Builder
	b.Write([]byte(v.Field))
	b.Write([]byte(":"))
	b.Write([]byte(v.Err.Error()))
	return b.String()
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	var b strings.Builder

	for i, err := range v {
		b.Write([]byte(err.Error()))
		if i != len(v)-1 {
			b.Write([]byte(" "))
		}
	}
	return b.String()
}

var ErrExpectedStruct = errors.New("expected struct")

func Validate(v interface{}) error {
	reflectV := reflect.ValueOf(v)
	typeV := reflectV.Type()
	validationErrors := make(ValidationErrors, 0)

	if typeV.Kind() != reflect.Struct {
		return ErrExpectedStruct
	}

	for i := 0; i < reflectV.NumField(); i++ {
		value := reflectV.Field(i)
		field := typeV.Field(i)
		tagV := field.Tag
		rules, err := SplitRules(tagV.Get(tagValue))
		if err != nil {
			log.Panic(err)
		}

		validErrors := validateField(field.Name, value, rules)
		if validErrors != nil {
			validationErrors = append(validationErrors, validErrors...)
		}
	}
	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateField(fieldName string, value reflect.Value, rules []Rule) ValidationErrors {
	var validationErrors ValidationErrors
	switch value.Kind() { //nolint:exhaustive
	case reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			err := validateField(fieldName, value.Index(i), rules)
			if err != nil {
				validationErrors = append(validationErrors, err...)
			}
		}
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

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
