package main

import (
	"fmt"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
}

func main() {
	// Отрезаем имя программы, оставляем только то, что ввел пользователь
	args := os.Args[1:]

	// Проверяем количество аргументов (от 1 до 3)
	if len(args) == 0 || len(args) > 3 {
		usage()
		return
	}

	var outputFile string
	var input string
	bannerName := "standard"

	// 1. Пытаемся найти флаг --output
	if strings.HasPrefix(args[0], "--output=") {
		outputFile = strings.TrimPrefix(args[0], "--output=")
		if outputFile == "" {
			usage()
			return
		}
		// Убираем обработанный флаг из списка
		args = args[1:]
	} else if strings.HasPrefix(args[0], "--") {
		usage()
		return
	}

	// 2. После флага первым делом должен идти текст
	if len(args) == 0 {
		usage()
		return
	}
	input = args[0]
	args = args[1:] // Убираем текст из списка

	// 3. Если что-то осталось — это название баннера
	if len(args) > 0 {
		bannerName = strings.TrimSuffix(args[0], ".txt")
	}

	// 4. Проверки и логика
	if input == "" {
		return
	}

	if !isASCII(input) {
		fmt.Println("error: non-ASCII symbol detected")
		return
	}

	// Загрузка и конвертация (эти функции должны быть в твоем проекте)
	if err := VerifyBannerHash(bannerName); err != nil {
		fmt.Println(err)
		return
	}

	banner, err := ReadBanner(bannerName)
	if err != nil {
		fmt.Println(err)
		return
	}

	output := ConvertToASCII(input, banner)

	// 5. Сохранение результата с теми самыми правами 0644
	if outputFile != "" {
		// Мы (6) пишем/читаем, остальные (4) только читают.
		err := os.WriteFile(outputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	fmt.Print(output)
}

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
