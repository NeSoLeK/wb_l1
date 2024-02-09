package main

import "fmt"

func t12() {
	// Создаем множество строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем пустое множество для хранения уникальных строк
	uniqueSet := make(map[string]bool)

	// Заполняем множество уникальными строками из последовательности
	for _, str := range sequence {
		uniqueSet[str] = true
	}

	// Выводим уникальные строки
	fmt.Println("Уникальные строки:")
	for str := range uniqueSet {
		fmt.Println(str)
	}
}
