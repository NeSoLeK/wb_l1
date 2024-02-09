package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(sentence string) string {
	// Разбиваем строку на слова
	words := strings.Fields(sentence)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Собираем слова обратно в строку
	reversedSentence := strings.Join(words, " ")

	return reversedSentence
}

func t20() {
	// Пример строки
	sentence := "snow dog sun"

	// Переворачиваем слова в строке
	reversed := reverseWords(sentence)

	fmt.Println("Перевернутые слова:", reversed)
}
