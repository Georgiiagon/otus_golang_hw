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
	panic("implement me")
}

func Validate(v interface{}) error {
	reflectV := reflect.ValueOf(v)
	typeV := reflectV.Type()
	for i := 0; i < reflectV.NumField(); i++ {
		value := reflectV.Field(i)
		t := typeV.Field(i)
		tagV := t.Tag
		fmt.Println(value, tagV.Get("validate"))
	}
	return nil
}
