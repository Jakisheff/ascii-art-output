package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"strings"
)

// expectedHashes содержит ожидаемые SHA-256 хэши для баннеров, чтобы уберечь их от изменения
var expectedHashes = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

func VerifyBannerHash(bannerName string) error {
	filename := bannerName + ".txt"
	expectedHash, ok := expectedHashes[filename]

	// Вычисляем хэш файла
	hash, err := GetFileHash("banners/" + filename)
	if err != nil {
		return fmt.Errorf("error calculating hash: %v", err)
	}

	// Сравниваем вычисленный хэш с ожидаемым
	if ok && hash != expectedHash {
		return fmt.Errorf("error: banner %s has been modified or corrupted (hash mismatch)", bannerName)
	}

	return nil
}

// GetFileHash читает файл и возвращает его SHA-256 хэш в виде строки
func GetFileHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("error hashing file: %v", err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil)), nil
}

// ReadBanner читает файл баннера и преобразует его в карту (map), где ключ - это rune,
// а значение - срез строк (массив высотой из 8 элементов)
func ReadBanner(name string) (map[rune][]string, error) {
	file, err := os.Open("banners/" + name + ".txt")
	if err != nil {
		return nil, fmt.Errorf("error: banner %s not found", name)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	banner := make(map[rune][]string)

	var char rune = 32 // Начинаем с пробела, код 32
	var charLines []string

	// Построчно читаем файл
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimRight(line, "\r") // Убираем возможный символ возврата каретки

		// Пропускаем начальные пустые строки между символами
		if line == "" && len(charLines) == 0 {
			continue
		}

		charLines = append(charLines, line)

		// Когда набралось 8 строк - сохраняем символ в карту
		if len(charLines) == 8 {
			banner[char] = charLines
			char++          // Переходим к следующему символу
			charLines = nil // Очищаем буфер для нового символа
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading banner file: %v", err)
	}

	return banner, nil
}
