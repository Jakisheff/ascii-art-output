package main

import (
	"strings"
)

// ConvertToASCII преобразует обычный текст в ASCII-арт, используя заданный баннер
func ConvertToASCII(text string, banner map[rune][]string) string {
	// Заменяем литеральные символы "\n" на фактические символы переноса строки
	text = strings.ReplaceAll(text, "\\n", "\n")
	if text == "" {
		return ""
	}

	// Проверяем, состоит ли текст только из переносов строк (например, "\n\n")
	isOnlyNewlines := true
	for _, ch := range text {
		if ch != '\n' {
			isOnlyNewlines = false
			break
		}
	}

	// Если текст состоит только из \n, просто выводим их
	if isOnlyNewlines {
		return text
	}

	// Разбиваем текст по переносам строк
	lines := strings.Split(text, "\n")
	var resultBuilder strings.Builder

	// Перебираем каждое слово/строку
	for _, line := range lines {
		// Если это пустая строка (например, из-за подряд идущих \n), то выводим перенос
		if line == "" {
			resultBuilder.WriteString("\n")
			continue
		}

		// У каждого символа ASCII-арта высота 8 строк, поэтому цикл идет до 8
		for i := 0; i < 8; i++ {
			// Проходим по каждому символу в слове
			for _, ch := range line {
				if asciiArr, ok := banner[ch]; ok {
					if i < len(asciiArr) {
						resultBuilder.WriteString(asciiArr[i]) // Добавляем i-ю строку символа
					}
				}
			}
			resultBuilder.WriteString("\n") // В конце каждой части символов переходим на новую строку
		}
	}

	return resultBuilder.String()
}
