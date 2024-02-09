package main

//Разработать программу, которая проверяет, что все символы в строке
//уникальные (true — если уникальные, false etc). Функция проверки должна быть
//регистронезависимой.

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов в строке
func isUnique(str string) bool {
	// Создаем мапу для отслеживания уникальных символов
	seen := make(map[rune]bool)

	// Проходим по каждому символу в строке
	for _, char := range strings.ToLower(str) {
		// Если символ уже встречался, возвращаем false
		if seen[char] {
			return false
		}
		// Помечаем символ как встреченный
		seen[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Примеры строк для проверки
	strings := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверяем каждую строку
	for _, str := range strings {
		fmt.Printf("%s - %t\n", str, isUnique(str))
	}
}
