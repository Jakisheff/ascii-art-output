package main

import (
	"fmt"
	"os"
	"strings"
)

// usage выводит инструкцию по использованию программы
func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func main() {
	// Проверка на минимальное и максимальное количество аргументов
	if len(os.Args) < 2 || len(os.Args) > 4 {
		usage()
		return
	}

	var outputFile string
	var input string
	bannerName := "standard"

	// Индекс, где находится введенный текст (по умолчанию 1)
	textIndex := 1

	// Проверяем, передан ли флаг --output
	firstArg := os.Args[1]
	if strings.HasPrefix(firstArg, "--output=") {
		outputFile = strings.TrimPrefix(firstArg, "--output=")
		if outputFile == "" {
			// Если после = ничего нет (просто --output=)
			usage()
			return
		}
		// Так как первый аргумент — это флаг --output, текст сдвигается на индекс 2
		textIndex = 2
	} else if strings.HasPrefix(firstArg, "--") {
		// Если это неизвестный флаг (например, --test)
		usage()
		return
	}

	// Если аргументов больше нет, но текст должен быть
	if textIndex >= len(os.Args) {
		usage()
		return
	}

	// Получаем введенный текст
	input = os.Args[textIndex]

	// Если есть следующий аргумент, то это шаблон (баннер)
	if textIndex+1 < len(os.Args) {
		bannerName = strings.TrimSuffix(os.Args[textIndex+1], ".txt") // Убираем .txt, если передали
	}

	// Если передали слишком много аргументов после текста и баннера
	if textIndex+2 < len(os.Args) {
		usage()
		return
	}

	// Если текст пустой, программа завершится
	if input == "" {
		return
	}

	// Проверяем, что текст содержит только допустимые ASCII символы
	if !isASCII(input) {
		fmt.Println("error: non-ASCII symbol detected")
		return
	}

	// Загружаем хэш файла и проверяем на изменения
	if err := VerifyBannerHash(bannerName); err != nil {
		fmt.Println(err)
		return
	}

	// Загружаем баннер
	banner, err := ReadBanner(bannerName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Выводим текст в виде ASCII-графики
	output := ConvertToASCII(input, banner)

	// Записываем результат: в файл или в консоль
	if outputFile != "" {
		err := os.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	fmt.Print(output)
}

// isASCII проверяет, состоит ли строка только из печатных символов ASCII
func isASCII(str string) bool {
	for _, char := range str {
		if char < 32 || char > 126 {
			if char != '\n' && char != '\r' {
				return false
			}
		}
	}
	return true
}
