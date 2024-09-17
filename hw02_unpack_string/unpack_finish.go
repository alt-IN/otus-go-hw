package main

import (
	"errors"
	"fmt"
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

func main() {
	var str = "45"
	fmt.Println(len(str))
	temp := []rune(str)
	for i, v := range temp {
		fmt.Println(i, v, string(v))
	}

	// str2 := strings.Builder{}
	// _, _ = str2.WriteString(str[0:3])
	// _, _ = str2.WriteString(strings.Repeat(string(str2.String()[2]), 0))
	str3, err := Unpack(str)
	fmt.Printf("%#v %#v", str3, err)
}
