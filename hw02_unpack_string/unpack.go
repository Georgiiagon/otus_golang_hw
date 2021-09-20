package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(initStr string) (string, error) {
	var tmpLetter rune
	var resultString []rune
	for i, letter := range initStr {
		if unicode.IsLetter(letter) || letter == 10 || (i != 0 && initStr[i-1] == 92 && tmpLetter != 92) {
			tmpLetter = letter
			resultString = append(resultString, tmpLetter)
			continue
		}

		if tmpLetter == 0 && (i != 0 || i == 0) && letter != 92 {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(letter) && letter != '0' {
			repeatTimes, _ := strconv.Atoi(string(letter))

			for i := repeatTimes; i > 1; i-- {
				resultString = append(resultString, tmpLetter)
			}

			tmpLetter = 0
			continue
		} else if letter == '0' {
			resultString = resultString[:len(resultString)-1]
		}
		tmpLetter = 0
	}

	return string(resultString), nil
}
