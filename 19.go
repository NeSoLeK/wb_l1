package main

import (
	"fmt"
)

// Функция для переворачивания строки с учетом символов Unicode
func reverseString(input string) string {
	// Преобразуем строку в руны (Unicode символы)
	runes := []rune(input)
	// Определяем длину строки
	length := len(runes)

	// Переворачиваем строку
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func t19() {
	// Пример строки
	input := "главрыба — абырвалг"

	// Переворачиваем строку
	reversed := reverseString(input)

	fmt.Println("Перевернутая строка:", reversed)
}
