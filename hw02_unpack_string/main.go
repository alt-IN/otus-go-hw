package hw02unpackstring

import (
	"fmt"
	"strconv"
	"unicode"
)

func repeatChars(input string) string {
	result := ""
	var currentRun string

	for i := 0; i < len(input); i++ {
		char := rune(input[i])
		if char == '\\' && i < len(input)-1 { // Проверка на экранирование
			i++                            // Переходим к следующему символу
			currentRun += string(input[i]) // Добавляем экранированный символ к текущей последовательности
			continue
		}

		if unicode.IsDigit(char) { // Если символ - цифра
			// Проверяем, что это первая цифра подряд, если есть предыдущий символ
			if i == 0 || !unicode.IsDigit(rune(input[i-1])) {
				if count, err := strconv.Atoi(string(char)); err == nil {
					// Добавляем текущую последовательность count раз
					result += currentRun
					for j := 1; j < count; j++ {
						result += currentRun
					}
				}
				currentRun = "" // Сбрасываем текущую последовательность
			}
		} else {
			// Добавляем символ в текущую последовательность
			currentRun += string(char)
		}
	}

	// Добавляем оставшуюся часть, если строка не заканчивается цифрой
	result += currentRun

	return result
}

func main() {
	input := "ab@3c#2d$\\4\\2"
	output := repeatChars(input)
	fmt.Println(output) // Вывод: "ab@ab@abcc#cc$$42"
}
