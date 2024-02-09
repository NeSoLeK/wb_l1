package main

//Реализовать пересечение двух неупорядоченных множеств.

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool) // Создаем пустое множество для результата

	// Перебираем элементы первого множества
	for key := range set1 {
		// Если элемент также есть во втором множестве, добавляем его в результат
		if set2[key] {
			result[key] = true
		}
	}

	return result
}

func t11() {
	// Пример множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Находим пересечение множеств
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Printf("Intersection: %v", result)
}
