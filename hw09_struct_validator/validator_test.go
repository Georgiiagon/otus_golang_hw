package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/pioz/faker"
	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
	}

	App struct {
		Number    int      `validate:"max:100|min:55"`
		Version   string   `validate:"len:5|in:glass,maxon|regexp:^g[lq]?a?ss"`
		testArray []string `validate:"len:1|in:a,b,c,d,e,z"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}

	WrongValidation struct {
		ID          int    `validate:"int:200"`
		Description string `validate:"test:100"`
	}

	WrongValidationSecond struct {
		ID          int    `validate:"max:a"`
		Description string `validate:"test:a"`
	}
)

func TestValidate(t *testing.T) {
	faker.SetSeed(time.Now().UnixNano())
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: ErrLen},
				ValidationError{Field: "Age", Err: ErrMin},
				ValidationError{Field: "Email", Err: ErrRegexp},
				ValidationError{Field: "Role", Err: ErrUnexpectedValue},
			},
		},
		{
			in: User{
				ID:     faker.DigitsWithSize(36),
				Name:   faker.Username(),
				Age:    faker.IntInRange(18, 50),
				Email:  faker.SafeEmail(),
				Role:   "admin",
				Phones: []string{faker.DigitsWithSize(11), faker.DigitsWithSize(11), "89222222222"},
				meta:   json.RawMessage("{\"test\": \"test\"}"),
			},
			expectedErr: nil,
		},

		{
			in: User{
				ID:     faker.DigitsWithSize(36),
				Name:   faker.Username(),
				Age:    faker.IntInRange(12, 17),
				Email:  "@example.com",
				Role:   "admin",
				Phones: []string{faker.DigitsWithSize(11), faker.DigitsWithSize(11), "8922222222"},
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Age", Err: ErrMin},
				ValidationError{Field: "Email", Err: ErrRegexp},
				ValidationError{Field: "Phones", Err: ErrLen},
			},
		},
		{
			in:          App{Number: 100, Version: "glass", testArray: []string{"a", "b", "c"}},
			expectedErr: nil,
		},
		{
			in: App{Number: 101, Version: "glass", testArray: []string{"a", "b", "c"}},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Number", Err: ErrMax},
			},
		},
		{
			in:          Token{Header: []byte("smth"), Payload: []byte("test-test"), Signature: []byte("signature")},
			expectedErr: nil,
		},
		{
			in:          Response{Code: 200, Body: "{\"status\": \"ok\"}"},
			expectedErr: nil,
		},
		{
			in:          Response{Code: 404, Body: "{\"status\": \"ok\"}"},
			expectedErr: nil,
		},
		{
			in: Response{Code: 504, Body: ""},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Code", Err: ErrUnexpectedValue},
			},
		},
		{
			in: WrongValidation{ID: faker.Int(), Description: faker.String()},
			expectedErr: ValidationErrors{
				ValidationError{Field: "ID", Err: ErrUnexpectedRule},
				ValidationError{Field: "Description", Err: ErrUnexpectedRule},
			},
		},
		{
			in:          "string",
			expectedErr: ErrExpectedStruct,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			if tt.expectedErr == nil {
				require.ErrorIs(t, err, tt.expectedErr)
			} else {
				require.EqualError(t, err, tt.expectedErr.Error())
			}
		})
	}
}

func TestOsExit(t *testing.T) {
	fakeExitOne := func() {
		Validate(WrongValidationSecond{Description: "test"})
	}
	fakeExitTwo := func() {
		Validate(WrongValidationSecond{ID: 1})
	}

	require.Panics(t, fakeExitOne)
	require.Panics(t, fakeExitTwo)
}
