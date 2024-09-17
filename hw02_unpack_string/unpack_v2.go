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
	escapeFlag := false
	for idx, val := range tmpRune {
		if escapeFlag == true {
			currElem = val
			escapeFlag = false
		} else {
			if num, err := strconv.Atoi(string(val)); err == nil {
				if idx == 0 || (unicode.IsDigit(tmpRune[idx-1]) && tmpRune[idx-1] != currElem) {
					return "", ErrInvalidString
				}
				_, _ = resultStr.WriteString(strings.Repeat(string(currElem), num))
			} else {
				if val == '\\' && idx < len(tmpRune) {
					escapeFlag = true
				} else {
					currElem = val
				}
			}
		}
	}
	fmt.Println(resultStr.String())
	return resultStr.String(), nil
}

func main() {
	var str = "ab6c\\3d\n4Я"
	fmt.Println(len(str))
	temp := []rune(str)
	for i, v := range temp {
		fmt.Println(i, v, string(v))
	}

	str2 := strings.Builder{}
	_, _ = str2.WriteString(str[0:3])
	_, _ = str2.WriteString(strings.Repeat(string(str2.String()[2]), 0))
	str3, _ := Unpack(str)
	fmt.Sprintf("%#v", str3)
}
