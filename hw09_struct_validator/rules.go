package hw09structvalidator

import (
	"errors"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	name  string
	value interface{}
}

var (
	ErrUnexpectedRule  = errors.New("rule is unexpected")
	ErrUnexpectedValue = errors.New("value is unexpected")
	ErrLen             = errors.New("length should be different")
	ErrRegexp          = errors.New("regexp is not matched")
	ErrMax             = errors.New("number is too big")
	ErrMin             = errors.New("number is too small")
)

func (r Rule) ValidateString(value reflect.Value) error {
	stringValue := value.String()
	switch r.name {
	case "in":
		for _, v := range r.value.([]string) {
			if stringValue == v {
				return nil
			}
		}
		return ErrUnexpectedValue

	case "len":
		if len(stringValue) == r.value.(int) {
			return nil
		}
		return ErrLen
	case "regexp":
		matched, err := regexp.MatchString(r.value.(string), stringValue)
		if err != nil {
			log.Panic(err)
		}

		if matched {
			return nil
		}

		return ErrRegexp
	}

	return ErrUnexpectedRule
}

func (r Rule) ValidateInt(value reflect.Value) error {
	intValue := int(value.Int())
	switch r.name {
	case "in":
		for _, val := range r.value.([]string) {
			expectedVal, err := strconv.Atoi(val)
			if err != nil {
				log.Panic(err)
			}

			if intValue == expectedVal {
				return nil
			}
		}

		return ErrUnexpectedValue

	case "min":
		if intValue >= r.value.(int) {
			return nil
		}
		return ErrMin
	case "max":
		if intValue <= r.value.(int) {
			return nil
		}
		return ErrMax
	}

	return ErrUnexpectedRule
}

func SplitRules(stringRules string) ([]Rule, error) {
	rulesArr := strings.Split(stringRules, "|")
	rules := make([]Rule, 0, len(rulesArr))
	var value interface{}
	var err error
	for _, r := range rulesArr {
		ruleArr := strings.Split(r, ":")
		if len(ruleArr) != 2 {
			continue
		}

		switch ruleArr[0] {
		case "in":
			value = strings.Split(ruleArr[1], ",")
		case "regexp":
			value = ruleArr[1]
		default:
			value, err = strconv.Atoi(ruleArr[1])
			if err != nil {
				return nil, err
			}
		}

		rules = append(rules, Rule{name: ruleArr[0], value: value})
	}

	return rules, nil
}
