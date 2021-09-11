package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(initStr string) (string, error) {
	var tmpLetter string
	var resultString string
	for i, letter := range initStr {
		if (letter >= 'a' && letter <= 'z') || letter == 10 || (i != 0 && initStr[i-1] == 92 && tmpLetter != `\`) {
			tmpLetter = string(letter)
			resultString += tmpLetter
			continue
		}

		if tmpLetter == "" && (i != 0 || i == 0) {
			return "", ErrInvalidString
		}

		if letter >= '1' && letter <= '9' {
			repeatTimes, _ := strconv.Atoi(string(letter))

			resultString += strings.Repeat(tmpLetter, repeatTimes-1)
			tmpLetter = ""
			continue
		} else if letter == '0' {
			resultString = resultString[:len(resultString)-1]
		}
		tmpLetter = ""
	}

	return resultString, nil
}
