package hw02unpackstring

import (
	"errors"
	"strconv"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(initStr string) (string, error) {
	var tmpLetter rune
	var resultString []rune
	for _, letter := range initStr {
		if (letter >= 'a' && letter <= 'z') || letter == 10 {
			tmpLetter = letter
			resultString = append(resultString, letter)
			continue
		}

		if tmpLetter == 0 {
			return "", ErrInvalidString
		}

		if letter >= '1' && letter <= '9' {
			repeatTimes, _ := strconv.Atoi(string(letter))
			for i := repeatTimes; i > 1; i-- {
				resultString = append(resultString, tmpLetter)
			}
			tmpLetter = 0
			continue
		}

		if letter == '0' {
			resultString = resultString[:len(resultString)-1]
		}
	}

	return string(resultString), nil
}
