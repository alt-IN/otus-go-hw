package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	// Place your code here.
	resultStr := strings.Builder{}
	tmpRune := []rune(str)
	var currElem rune
	for idx, val := range tmpRune {
		if num, err := strconv.Atoi(string(val)); err == nil {
			if idx == 0 || unicode.IsDigit(tmpRune[idx-1]) {
				return "", ErrInvalidString
			}
			_, _ = resultStr.WriteString(strings.Repeat(string(currElem), num))
		} else {
			if idx == len(tmpRune)-1 || !unicode.IsDigit((tmpRune[idx+1])) {
				_, _ = resultStr.WriteRune(val)
			}
			currElem = val
		}
	}
	return resultStr.String(), nil
}
