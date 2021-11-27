package hw09structvalidator

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field string
	Err   error
}

var (
	unexpectedRuleError = errors.New("Rule is unexpected!")
	inError             = errors.New("Value is unexpected!")
	lenError            = errors.New("Length should be different!")
	regexpError         = errors.New("regexp is not sutisfied!")
)

type Rule struct {
	name  string
	value interface{}
}

func (r Rule) ValidateString(value reflect.Value) error {
	stringValue := value.String()
	switch r.name {
	case "in":
		for _, v := range r.value.([]string) {
			if stringValue == v {
				return nil
			}
		}
		return inError

	case "len":
		if len(stringValue) == r.value.(int) {
			return nil
		}
		return lenError
	case "regexp":
		matched, err := regexp.MatchString(r.value.(string), stringValue)

		if err != nil {
			return err
		}

		if matched {
			return nil
		}

		return regexpError
	}

	return unexpectedRuleError
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
		rules := splitRules(tagV.Get("validate"))
		validateField(field.Name, value, rules)
	}
	fmt.Println(len(validationErrors), validationErrors)
	return nil
}

func validateField(fieldName string, value reflect.Value, rules []Rule) {
	switch value.Kind() {
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
		//...
	}

	fmt.Println(value, value.Kind(), value.Type(), rules)
	_ = rules
}

// func validateInt(value reflect.Value, rules []Rule) {
// 	for _, rule := range rules {

// 	}
// }

func validateString(fieldName string, value reflect.Value, rules []Rule) {
	for _, rule := range rules {
		err := rule.ValidateString(value)

		if err != nil {
			validationErrors = append(validationErrors, ValidationError{Field: fieldName, Err: err})
		}
	}
}

func splitRules(stringRules string) []Rule {
	rulesArr := strings.Split(stringRules, "|")
	rules := make([]Rule, 0, len(rulesArr))
	var value interface{}
	var err error
	for _, r := range rulesArr {
		ruleArr := strings.Split(r, ":")

		switch ruleArr[0] {
		case "in":
			value = strings.Split(ruleArr[1], ",")
		case "regexp":
			value = ruleArr[1]
		default:
			value, err = strconv.Atoi(ruleArr[1])
			if err != nil {
				log.Fatal(err)
			}
		}

		rules = append(rules, Rule{name: ruleArr[0], value: value})
	}

	return rules
}
